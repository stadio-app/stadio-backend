package services

import (
	"context"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/op/go-logging"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var log = logging.MustGetLogger("services")

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
			Email:        input.Email,
			Name:         input.Name,
			Password:     &hashed_password,
			AuthPlatform: model.UserAuthPlatformType_Internal,
			Active:       false,
			PhoneNumber:  input.PhoneNumber,
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
			Email:        input.Email,
			Name:         input.Name,
			AuthPlatform: oauth_type,
			Active:       true,
			PhoneNumber:  input.PhoneNumber,
		}).
		RETURNING(table.User.AllColumns)
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &user); err != nil {
		return gmodel.User{}, fmt.Errorf("user entry could not be created. %s", err.Error())
	}

	var activationCode string
	db := service.DbOrTxQueryable()
	codeQuery := table.EmailVerification.
		SELECT(table.EmailVerification.Code).
		WHERE(
			table.EmailVerification.UserID.EQ(
				postgres.Int(user.ID),
			),
		)
	if err := codeQuery.QueryContext(ctx, db, &activationCode); err != nil {
		return gmodel.User{}, fmt.Errorf("failed to retrieve activation code")
	}
	// TODO: send activation code to email
	log.Debugf("activation code: ", activationCode)

	return user, nil
}

func (service Service) VerifyUser(ctx context.Context, email string, activation_code string) (bool, error) {
	db := service.DbOrTxQueryable()
	query := table.User.
		SELECT(
			table.User.Active,
			table.EmailVerification.Code,
		).
		FROM(
			table.User.LEFT_JOIN(
				table.EmailVerification,
				table.User.ID.EQ(table.EmailVerification.UserID),
			)).
		WHERE(table.User.Email.EQ(postgres.String(email))).
		// TODO: There should be a unique constraint on email
		// so that this never returns more than one row.
		LIMIT(1)

	log.Debug(query.Sql())
	var verifyable_user gmodel.VerifyableUser
	if err := query.QueryContext(ctx, db, &verifyable_user); err != nil {
		return false, fmt.Errorf("incorrect email or password")
	}
	if verifyable_user.Active {
		return false, fmt.Errorf("user is already active")
	}
	log.Debug(activation_code)
	log.Debug(verifyable_user)
	log.Debug(verifyable_user.Code)
	if activation_code != verifyable_user.Code {
		return false, fmt.Errorf("invalid activation code")
	}

	var active_set bool
	update_query := table.User.
		UPDATE(table.User.Active).
		SET(true).
		WHERE(table.User.Email.EQ(postgres.String(email))).
		RETURNING(table.User.Active)
	if err := update_query.QueryContext(ctx, db, &active_set); err != nil {
		return false, fmt.Errorf("email could not be verified %s", err.Error())
	}

	return active_set, nil
}

func (service Service) GoogleAuthentication(ctx context.Context, access_token string) (gmodel.Auth, error) {
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
			Name:  userinfo.Name,
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
	/*
	* (Suggestion) TODO:
	* 1. Return the user even if they are inactive
	* 2. Check if the user is active or not on the client-side
	* 3. Make the email verification endpoint authenticatable (more secure).
	 */
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
		User:  &user,
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
