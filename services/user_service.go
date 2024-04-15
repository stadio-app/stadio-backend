package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
	"google.golang.org/api/oauth2/v2"
)

func (service Service) CreateInternalUser(ctx context.Context, input gmodel.CreateAccountInput) (gmodel.User, error) {
	if service.UserEmailExists(ctx, input.Email) {
		return gmodel.User{}, fmt.Errorf("email already exists")
	}
	hashed_password, hash_err := service.HashPassword(input.Password)
	if hash_err != nil {
		return gmodel.User{}, hash_err
	}

	// Create transaction
	tx, tx_err := service.DB.BeginTx(ctx, nil)
	if tx_err != nil {
		return gmodel.User{}, tx_err
	}

	var user gmodel.User
	qb := table.User.
		INSERT(
			table.User.Email,
			table.User.Name,
			table.User.Password,
			table.User.AuthPlatform,
			table.User.Active,
			table.User.PhoneNumber,
		).
		MODEL(model.User{
			Email: input.Email,
			Name: input.Name,
			Password: &hashed_password,
			AuthPlatform: model.UserAuthPlatformType_Internal,
			Active: false,
			PhoneNumber: input.PhoneNumber,
		}).
		RETURNING(table.User.AllColumns)
	if err := qb.QueryContext(ctx, tx, &user); err != nil {
		tx.Rollback()
		return gmodel.User{}, fmt.Errorf("user entry could not be created. %s", err.Error())
	}
	service.TX = tx
	if _, err := service.CreateEmailVerification(ctx, user); err != nil {
		tx.Rollback()
		return gmodel.User{}, err
	}

	// Commit changes from transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return gmodel.User{}, err
	}
	return user, nil
}

func (service Service) CreateOauthUser(ctx context.Context, input gmodel.CreateAccountInput, oauth_type model.UserAuthPlatformType) (gmodel.User, error) {
	var user gmodel.User
	qb := table.User.
		INSERT(
			table.User.Email,
			table.User.Name,
			table.User.AuthPlatform,
			table.User.Active,
			table.User.PhoneNumber,
		).
		MODEL(model.User{
			Email: input.Email,
			Name: input.Name,
			AuthPlatform: oauth_type,
			Active: true,
			PhoneNumber: input.PhoneNumber,
		}).
		RETURNING(table.User.AllColumns)
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &user); err != nil {
		return gmodel.User{}, fmt.Errorf("user entry could not be created. %s", err.Error())
	}
	return user, nil
}

func (service Service) GoogleAuthentication(ctx context.Context, access_token string) (gmodel.Auth, error) {
	res, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", access_token))
	if err != nil || res.StatusCode == http.StatusUnauthorized {
		return gmodel.Auth{}, fmt.Errorf("invalid access token")
	}
	userDataRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return gmodel.Auth{}, fmt.Errorf("could not read response body")
	}
	var userData oauth2.Userinfo
	if err := json.Unmarshal(userDataRaw, &userData); err != nil {
		return gmodel.Auth{}, err
	}

	var user gmodel.User
	new_user := false
	if service.UserEmailExists(ctx, userData.Email) {
		user, _ = service.FindUserByEmail(ctx, userData.Email)
	} else {
		user, err = service.CreateOauthUser(ctx, gmodel.CreateAccountInput{
			Email: userData.Email,
			Name: userData.Name,
		}, model.UserAuthPlatformType_Google)
		if err != nil {
			return gmodel.Auth{}, err
		}
		new_user = true
	}
	auth_state, err := service.CreateAuthStateWithJwt(ctx, user.ID)
	if err == nil && new_user {
		auth_state.IsNewUser = &new_user
	}
	return auth_state, err
}

func (service Service) LoginInternal(ctx context.Context, email string, password string) (gmodel.Auth, error) {
	db := service.DbOrTxQueryable()	
	query := table.User.
		SELECT(table.User.AllColumns).
		WHERE(table.User.Email.EQ(postgres.String(email))).
		LIMIT(1)
	var verify_user model.User
	if err := query.QueryContext(ctx, db, &verify_user); err != nil {
		return gmodel.Auth{}, fmt.Errorf("incorrect email or password")
	}
	if !verify_user.Active {
		return gmodel.Auth{}, fmt.Errorf("please verify your email")
	}
	if verify_user.Password == nil {
		return gmodel.Auth{}, fmt.Errorf("password has not been set for this account. try a different authentication method")
	}
	if !service.VerifyPasswordHash(password, *verify_user.Password) {
		return gmodel.Auth{}, fmt.Errorf("incorrect email or password")
	}

	return service.CreateAuthStateWithJwt(ctx, verify_user.ID)
}

func (service Service) CreateAuthStateWithJwt(ctx context.Context, user_id int64) (gmodel.Auth, error) {
	auth_state, auth_state_err := service.CreateAuthState(ctx, gmodel.User{
		ID: user_id,
	}, nil)
	if auth_state_err != nil {
		return gmodel.Auth{}, fmt.Errorf("could not create auth state")
	}

	qb := table.User.
		SELECT(table.User.AllColumns, table.AuthState.ID).
		FROM(table.User.LEFT_JOIN(
			table.AuthState, 
			table.User.ID.EQ(table.AuthState.UserID),
		)).
		WHERE(postgres.AND(
			table.User.ID.EQ(postgres.Int(user_id)),
			table.AuthState.ID.EQ(postgres.Int(auth_state.ID)),
		)).LIMIT(1)
	var user gmodel.User
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &user); err != nil {
		return gmodel.Auth{}, fmt.Errorf("internal error")
	}

	// Generate JWT
	jwt, err := service.GenerateJWT(service.Tokens.JwtKey, &user)
	if err != nil {
		return gmodel.Auth{}, fmt.Errorf("could not generate JWT")
	}
	return gmodel.Auth{
		Token: jwt,
		User: &user,
	}, nil
}

// Returns `false` if user email does not exist. Otherwise `true`
func (service Service) UserEmailExists(ctx context.Context, email string) bool {
	query := table.User.
		SELECT(table.User.Email.AS("email")).
		FROM(table.User).
		WHERE(table.User.Email.EQ(postgres.String(email))).
		LIMIT(1)
	var dest struct{ Email string }
	err := query.QueryContext(ctx, service.DbOrTxQueryable(), &dest)
	return err == nil
}

func (service Service) FindUserByEmail(ctx context.Context, email string) (gmodel.User, error) {
	qb := table.User.
		SELECT(table.User.AllColumns).
		WHERE(table.User.Email.EQ(postgres.String(email))).
		LIMIT(1)
	var user gmodel.User
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &user); err != nil {
		return gmodel.User{}, err
	}
	return user, nil
}
