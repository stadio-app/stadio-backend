package services

import (
	"context"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
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

	query := table.User.INSERT(
		table.User.Email,
		table.User.Name,
		table.User.Password,
		table.User.AuthPlatform,
		table.User.Active,
	).VALUES(
		input.Email,
		input.Name,
		hashed_password,
		model.UserAuthPlatformType_Internal,
		false,
	).RETURNING(table.User.AllColumns)

	var user gmodel.User
	err := query.QueryContext(ctx, tx, &user)
	if err != nil {
		tx.Rollback()
		return gmodel.User{}, fmt.Errorf("user entry could not be created")
	}
	if _, err := service.CreateEmailVerification(ctx, user, tx); err != nil {
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

func (service Service) LoginInternal(ctx context.Context, email string, password string) (gmodel.Auth, error) {
	query := table.User.
		SELECT(table.User.AllColumns).
		WHERE(table.User.Email.EQ(postgres.String(email))).
		LIMIT(1)
	var verify_user model.User
	if err := query.QueryContext(ctx, service.DB, &verify_user); err != nil {
		return gmodel.Auth{}, fmt.Errorf("incorrect email or password")
	}

	if !verify_user.Active {
		return gmodel.Auth{}, fmt.Errorf("please verify your email")
	}
	if verify_user.AuthPlatform != model.UserAuthPlatformType_Internal || verify_user.Password == nil {
		return gmodel.Auth{}, fmt.Errorf("login platform is %s", verify_user.AuthPlatform.String())
	}
	if !service.VerifyPasswordHash(password, *verify_user.Password) {
		return gmodel.Auth{}, fmt.Errorf("incorrect email or password")
	}

	auth_state, auth_state_err := service.CreateAuthState(ctx, gmodel.User{
		ID: verify_user.ID,
	}, nil)
	if auth_state_err != nil {
		return gmodel.Auth{}, fmt.Errorf("could not create auth state")
	}

	query = table.User.
		SELECT(
			table.User.AllColumns,
			table.AuthState.ID,
		).
		FROM(
			table.User.
				LEFT_JOIN(
					table.AuthState, 
					table.User.ID.EQ(table.AuthState.UserID),
				),
		).
		WHERE(
			table.User.ID.
				EQ(postgres.Int64(verify_user.ID)).
				AND(table.AuthState.ID.EQ(postgres.Int64(auth_state.ID))),
		).
		LIMIT(1)
	var user gmodel.User
	if err := query.QueryContext(ctx, service.DB, &user); err != nil {
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
		SELECT(table.User.Email).
		FROM(table.User).
		WHERE(
			table.User.Email.EQ(postgres.String(email)),
		).LIMIT(1)
	var dest struct{
		Email string
	}
	err := query.QueryContext(ctx, service.DB, &dest)
	if err != nil || dest.Email == "" {
		return false
	}
	return true
}
