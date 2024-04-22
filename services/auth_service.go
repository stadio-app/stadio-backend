package services

import (
	"context"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/golang-jwt/jwt"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

const EMAIL_VERIFICATION_CODE_LEN = 10

// Given an existing user, create an auth_state row
func (service Service) CreateAuthState(ctx context.Context, user gmodel.User, ip_address *string) (model.AuthState, error) {
	query := table.AuthState.INSERT(
		table.AuthState.UserID,
		table.AuthState.IPAddress,
	).MODEL(model.AuthState{
		UserID: user.ID,
		IPAddress: ip_address,
	}).RETURNING(table.AuthState.AllColumns)

	var auth_state model.AuthState
	err := query.QueryContext(ctx, service.DbOrTxQueryable(), &auth_state)
	return auth_state, err
}

func (service Service) CreateEmailVerification(ctx context.Context, user gmodel.User) (model.EmailVerification, error) {
	code, code_err := utils.GenerateRandomUrlEncodedString(EMAIL_VERIFICATION_CODE_LEN)
	if code_err != nil {
		return model.EmailVerification{}, code_err
	}

	query := table.EmailVerification.INSERT(
		table.EmailVerification.UserID,
		table.EmailVerification.Code,
	).MODEL(model.EmailVerification{
		UserID: user.ID,
		Code: code,
	}).RETURNING(table.EmailVerification.AllColumns)

	var email_verification model.EmailVerification
	err := query.QueryContext(ctx, service.DbOrTxQueryable(), &email_verification)
	return email_verification, err
}

func (service Service) ResendEmailVerification(ctx context.Context, user gmodel.User) (email_verification model.EmailVerification, err error) {
	if user.Active {
		return model.EmailVerification{}, fmt.Errorf("user already has a verified email address")
	}
	service.TX, err = service.DB.BeginTx(ctx, nil)
	if err != nil {
		return model.EmailVerification{}, err
	}

	_, err = table.EmailVerification.DELETE().
		WHERE(table.EmailVerification.UserID.EQ(postgres.Int(user.ID))).
		ExecContext(ctx, service.TX)
	if err != nil {
		service.TX.Rollback()
		return model.EmailVerification{}, fmt.Errorf("user email verification entry deletion failed")
	}

	email_verification, err = service.CreateEmailVerification(ctx, user)
	if err != nil {
		service.TX.Rollback()
		return model.EmailVerification{}, err
	}
	if err := service.TX.Commit(); err != nil {
		return model.EmailVerification{}, fmt.Errorf("could not commit changes")
	}
	return email_verification, nil
}

func (service Service) FindEmailVerificationByCode(ctx context.Context, verification_code string) (model.EmailVerification, error) {
	qb := table.EmailVerification.
		SELECT(table.EmailVerification.AllColumns).
		WHERE(table.EmailVerification.Code.EQ(postgres.String(verification_code))).
		LIMIT(1)
	var email_verification model.EmailVerification
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &email_verification); err != nil {
		return model.EmailVerification{}, fmt.Errorf("invalid email verification code")
	}
	return email_verification, nil
}

func (service Service) VerifyUserEmail(ctx context.Context, verification_code string) (gmodel.User, error) {
	var err error
	service.TX, err = service.DB.BeginTx(ctx, nil)
	if err != nil {
		service.TX.Rollback()
		return gmodel.User{}, err
	}

	email_verification, err := service.FindEmailVerificationByCode(ctx, verification_code)
	if err != nil {
		service.TX.Rollback()
		return gmodel.User{}, err
	}

	if time.Until(email_verification.CreatedAt).Abs() > time.Hour {
		service.TX.Rollback()
		// Delete verification entry since it's expired
		del_query := table.EmailVerification.
			DELETE().
			WHERE(table.EmailVerification.ID.EQ(postgres.Int(email_verification.ID)))
		if _, err := del_query.ExecContext(ctx, service.DB); err != nil {
			return gmodel.User{}, err
		}
		return gmodel.User{}, fmt.Errorf("verification code has expired")
	}

	update := table.User.
		UPDATE(table.User.Active, table.User.UpdatedAt).
		SET(postgres.Bool(true), postgres.DateT(time.Now())).
		WHERE(table.User.ID.EQ(postgres.Int(email_verification.UserID)))
	if _, err := update.ExecContext(ctx, service.TX); err != nil {
		service.TX.Rollback()
		return gmodel.User{}, fmt.Errorf("could not update user email verification status to verified")
	}

	// Remove email_verification row
	delete := table.EmailVerification.
		DELETE().
		WHERE(postgres.AND(
			table.EmailVerification.ID.EQ(postgres.Int(email_verification.ID)),
			table.EmailVerification.Code.EQ(postgres.String(verification_code)),
		))
	if _, err := delete.ExecContext(ctx, service.TX); err != nil {
		service.TX.Rollback()
		return gmodel.User{}, fmt.Errorf("could not delete email verification entry")
	}

	if err := service.TX.Commit(); err != nil {
		return gmodel.User{}, fmt.Errorf("could not commit changes")
	}
	service.TX = nil
	return service.FindUserById(ctx, email_verification.UserID)
}

func (Service) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func (Service) VerifyPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// Generates a JWT with claims, signed with key
func (Service) GenerateJWT(key string, user *gmodel.User) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		"authPlatform": user.AuthPlatform.String(),
		"authStateId": user.AuthStateID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	token, err := jwt.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service Service) VerifyJwt(ctx context.Context, authorization types.AuthorizationKeyType) (gmodel.User, error) {
	jwt_raw, err := authorization.GetToken()
	if err != nil {
		return gmodel.User{}, err
	}

	claims, err := utils.GetJwtClaims(jwt_raw, service.Tokens.JwtKey)
	if err != nil {
		return gmodel.User{}, err
	}
	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return gmodel.User{}, fmt.Errorf("token expired")
	}

	authStateId := int64(claims["authStateId"].(float64))
	userId := int64(claims["id"].(float64))
	email := claims["email"].(string)
	query := table.User.
		SELECT(
			table.User.AllColumns,
			table.AuthState.ID,
		).
		FROM(table.User.LEFT_JOIN(
			table.AuthState, 
			table.User.ID.EQ(table.AuthState.UserID),
		)).
		WHERE(
			table.User.ID.
				EQ(postgres.Int64(userId)).
				AND(table.User.Email.EQ(postgres.String(email))).
				AND(table.AuthState.ID.EQ(postgres.Int64(authStateId))),
		).
		LIMIT(1)
	var user gmodel.User
	if err := query.QueryContext(ctx, service.DbOrTxQueryable(), &user); err != nil {
		return gmodel.User{}, fmt.Errorf("one or more invalid claim values")
	}
	return user, nil
}

func (Service) GetAuthUserFromContext(ctx context.Context) gmodel.User {
	return ctx.Value(types.AuthUserKey).(gmodel.User)
}
