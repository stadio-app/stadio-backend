package services

import (
	"context"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
	"golang.org/x/crypto/bcrypt"
)

func (service Service) CreateInternalUser(ctx context.Context, input gmodel.CreateAccountInput) (gmodel.User, error) {
	if service.UserEmailExists(ctx, input.Email) {
		return gmodel.User{}, fmt.Errorf("email already exists")
	}
	hashed_password, hash_err := service.HashPassword(input.Password)
	if hash_err != nil {
		return gmodel.User{}, hash_err
	}

	query := table.User.INSERT(
		table.User.Email,
		table.User.Name,
		table.User.Password,
		table.User.AuthPlatform,
	).VALUES(
		input.Email,
		input.Name,
		hashed_password,
		model.UserAuthPlatformType_Internal,
	).RETURNING(table.User.AllColumns)

	var user gmodel.User
	err := query.QueryContext(ctx, service.DB, &user)
	return user, err
}

// Returns `false` if user email does not exist. Otherwise `true`
func (service Service) UserEmailExists(ctx context.Context, email string) bool {
	query := table.User.
		SELECT(table.User.Email).
		FROM(table.User).
		WHERE(
			table.User.Email.EQ(postgres.String(email)),
		).LIMIT(1)
	var dest model.User
	err := query.QueryContext(ctx, service.DB, &dest)
	if err != nil || dest.Email == "" {
		return false
	}
	return true
}

func (Service) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func (Service) VerifyPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
