package services

import (
	"context"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func (service Service) CreateInternalUser(ctx context.Context, input gmodel.CreateAccountInput) (user gmodel.User, email_verification model.EmailVerification, err error) {
	if service.UserEmailExists(ctx, input.Email) {
		return gmodel.User{}, model.EmailVerification{}, fmt.Errorf("email already exists")
	}
	hashed_password, hash_err := service.HashPassword(input.Password)
	if hash_err != nil {
		return gmodel.User{}, model.EmailVerification{}, hash_err
	}

	// Create transaction
	tx, tx_err := service.DB.BeginTx(ctx, nil)
	if tx_err != nil {
		return gmodel.User{}, model.EmailVerification{}, tx_err
	}

	qb := table.User.
		INSERT(
			table.User.Email,
			table.User.Name,
			table.User.Password,
			table.User.Active,
			table.User.PhoneNumber,
		).
		MODEL(model.User{
			Email: input.Email,
			Name: input.Name,
			Password: &hashed_password,
			Active: false,
			PhoneNumber: input.PhoneNumber,
		}).
		RETURNING(table.User.AllColumns)
	if err := qb.QueryContext(ctx, tx, &user); err != nil {
		tx.Rollback()
		return gmodel.User{}, model.EmailVerification{}, fmt.Errorf("user entry could not be created. %s", err.Error())
	}
	service.TX = tx
	if email_verification, err = service.CreateEmailVerification(ctx, user); err != nil {
		tx.Rollback()
		return gmodel.User{}, model.EmailVerification{}, err
	}

	// Commit changes from transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return gmodel.User{}, model.EmailVerification{}, err
	}
	return user, email_verification, nil
}

func (service Service) CreateOauthUser(ctx context.Context, input gmodel.CreateAccountInput, oauth_type model.UserAuthPlatformType) (gmodel.User, error) {
	var user gmodel.User
	qb := table.User.
		INSERT(
			table.User.Email,
			table.User.Name,
			table.User.Active,
			table.User.PhoneNumber,
		).
		MODEL(model.User{
			Email: input.Email,
			Name: input.Name,
			Active: true,
			PhoneNumber: input.PhoneNumber,
		}).
		RETURNING(table.User.AllColumns)
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &user); err != nil {
		return gmodel.User{}, fmt.Errorf("user entry could not be created. %s", err.Error())
	}
	return user, nil
}

func (service Service) GoogleAuthentication(ctx context.Context, access_token string, ip_address *string) (gmodel.Auth, error) {
	oauth_service, err := oauth2.NewService(ctx, option.WithoutAuthentication())
	if err != nil {
		return gmodel.Auth{}, fmt.Errorf("could not create service")
	}
	userinfo_service := oauth2.NewUserinfoService(oauth_service)
	userinfo, err := userinfo_service.Get().Do(googleapi.QueryParameter("access_token", access_token))
	if err != nil {
		return gmodel.Auth{}, fmt.Errorf("invalid access token")
	}

	var user gmodel.User
	new_user := false
	if service.UserEmailExists(ctx, userinfo.Email) {
		user, _ = service.FindUserByEmail(ctx, userinfo.Email)
	} else {
		user, err = service.CreateOauthUser(ctx, gmodel.CreateAccountInput{
			Email: userinfo.Email,
			Name: userinfo.Name,
		}, model.UserAuthPlatformType_Google)
		if err != nil {
			return gmodel.Auth{}, err
		}
		new_user = true
	}
	auth_state, err := service.CreateAuthStateWithJwt(ctx, user.ID, model.UserAuthPlatformType_Google, ip_address)
	if err == nil && new_user {
		auth_state.IsNewUser = &new_user
	}
	return auth_state, err
}

func (service Service) LoginInternal(ctx context.Context, email string, password string, ip_address *string) (gmodel.Auth, error) {
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

	return service.CreateAuthStateWithJwt(ctx, verify_user.ID, model.UserAuthPlatformType_Internal, ip_address)
}

func (service Service) CreateAuthStateWithJwt(
	ctx context.Context, 
	user_id int64, 
	auth_platform model.UserAuthPlatformType, 
	ip_address *string,
) (gmodel.Auth, error) {
	auth_state, auth_state_err := service.CreateAuthState(ctx, gmodel.User{ ID: user_id }, auth_platform, ip_address)
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

func (service Service) FindUserById(ctx context.Context, id int64) (gmodel.User, error) {
	qb := table.User.
		SELECT(table.User.AllColumns).
		WHERE(table.User.ID.EQ(postgres.Int(id))).
		LIMIT(1)
	var user gmodel.User
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &user); err != nil {
		return gmodel.User{}, err
	}
	return user, nil
}
