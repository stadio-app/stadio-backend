package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/golang-jwt/jwt"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
	"github.com/stadio-app/stadio-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

const EMAIL_VERIFICATION_CODE_LEN = 10

// Given an existing user (optional)
func (service Service) CreateAuthState(ctx context.Context, user gmodel.User, ip_address *string) (model.AuthState, error) {
	query := table.AuthState.INSERT(
		table.AuthState.UserID,
		table.AuthState.IPAddress,
	).VALUES(
		user.ID,
		ip_address,
	).RETURNING(table.AuthState.AllColumns)

	var auth_state model.AuthState
	err := query.QueryContext(ctx, service.DB, &auth_state)
	return auth_state, err
}

func (service Service) CreateEmailVerification(ctx context.Context, user gmodel.User, tx *sql.Tx) (model.EmailVerification, error) {
	var db qrm.Queryable = service.DB
	if tx != nil {
		db = tx
	}
	code, code_err := utils.GenerateRandomUrlEncodedString(EMAIL_VERIFICATION_CODE_LEN)
	if code_err != nil {
		return model.EmailVerification{}, code_err
	}

	query := table.EmailVerification.INSERT(
		table.EmailVerification.UserID,
		table.EmailVerification.Code,
	).VALUES(
		user.ID,
		code,
	).RETURNING(table.EmailVerification.AllColumns)

	var email_verification model.EmailVerification
	err := query.QueryContext(ctx, db, &email_verification)
	return email_verification, err
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
func (s *Service) GenerateJWT(key string, user *gmodel.User) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		"authPlatform": user.AuthPlatform.String(),
		"authStateId": user.AuthStateID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	token, err := jwt.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}
