package tests

import (
	"testing"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
	"github.com/stadio-app/stadio-backend/utils"
)

func TestUser(t *testing.T) {
	var user1 gmodel.User
	user1_input := gmodel.CreateAccountInput{
		Email: "user1@email.com",
		Name: "User 1",
		Password: "password123",
	}
	var (
		user1_auth gmodel.Auth
		user1_auth2 gmodel.Auth
	)
	
	t.Run("create user", func(t *testing.T) {
		var err error
		user1, err = service.CreateInternalUser(ctx, user1_input)
		if err != nil {
			t.Fatal(err)
		}
		if user1.Active {
			t.Fatal("User.Active should be set to false")
		}
		if user1.AuthPlatform != gmodel.AuthPlatformTypeInternal {
			t.Fatal("user auth platform should be internal")
		}

		t.Parallel()

		t.Run("password must be hashed", func(t *testing.T) {
			var u model.User
			qb := table.User.SELECT(table.User.AllColumns).
				FROM(table.User).
				WHERE(table.User.ID.EQ(postgres.Int(user1.ID))).
				LIMIT(1)
			if err := qb.QueryContext(ctx, db, &u); err != nil {
				t.Fatal(err)
			}

			if *u.Password == user1_input.Password {
				t.Fatal("password should not be stored in plain text")
			}

			new_hash_pass, err := service.HashPassword(user1_input.Password)
			if err != nil {
				t.Fatal("could not hash password", err)
			}
			if *u.Password == new_hash_pass {
				t.Fatal("hashes of the same password should not result in the same hash")
			}
		})

		t.Run("should create email verification entry", func(t *testing.T) {
			qb := table.EmailVerification.
				SELECT(table.EmailVerification.AllColumns).
				FROM(table.EmailVerification).
				WHERE(table.EmailVerification.UserID.EQ(postgres.Int(user1.ID))).
				LIMIT(1)
			var email_verification model.EmailVerification
			if err := qb.QueryContext(ctx, db, &email_verification); err != nil {
				t.Fatal("email verification entry was not created.", err.Error())
			}
		})

		t.Run("duplicate user", func(t *testing.T) {
			_, err := service.CreateInternalUser(ctx, gmodel.CreateAccountInput{
				Email: user1.Email,
				Password: "abc123",
			})
	
			if err == nil {
				t.Fatal("user email is duplicate. should not create user.")
			}
		})
	})

	t.Run("login", func(t *testing.T) {
		t.Run("correct input", func(t *testing.T) {
			t.Run("inactive user", func(t *testing.T) {
				_, err := service.LoginInternal(ctx, user1_input.Email, user1_input.Password)
				if err == nil {
					t.Fatal("should not login. user is still inactive")
				}
			})

			t.Run("active user", func(t *testing.T) {
				var err error
				// activate user account first...
				_, err = table.User.UPDATE(table.User.Active).
					SET(postgres.Bool(true)).
					WHERE(table.User.ID.EQ(postgres.Int(user1.ID))).
					ExecContext(ctx, db)
				if err != nil {
					t.Fatal("could not update user active status", err.Error())
				}
				user1.Active = true

				user1_auth, err = service.LoginInternal(ctx, user1_input.Email, user1_input.Password)
				if err != nil {
					t.Fatal("could not login", err.Error())
				}
				user1.AuthStateID = user1_auth.User.AuthStateID
				if *user1_auth.User != user1 {
					t.Fatal("returned user object does not match")
				}

				t.Parallel()
				
				t.Run("verify jwt", func(t *testing.T) {
					claims, err := utils.GetJwtClaims(user1_auth.Token, app.Tokens.JwtKey)
					if err != nil {
						t.Fatal("invalid jwt", err.Error())
					}

					claims_user_id, ok := claims["id"].(float64)
					if !ok {
						t.Fatal("could not convert claims.id to float64")
					}
					if int64(claims_user_id) != user1.ID {
						t.Fatal("jwt claim user.id does not match")
					}
				})

				t.Run("check if auth_state entry exists", func(t *testing.T) {
					var auth_state model.AuthState
					qb := table.AuthState.SELECT(table.AuthState.AllColumns).
						FROM(table.AuthState).
						WHERE(table.AuthState.ID.EQ(postgres.Int(*user1_auth.User.AuthStateID))).
						LIMIT(1)
					if err := qb.QueryContext(ctx, db, &auth_state); err != nil {
						t.Fatal("could not find auth_state entry", err.Error())
					}

					if auth_state.UserID != user1.ID {
						t.Fatal("auth_state.user_id does not match")
					}
				})

				t.Run("new login should create new auth_state entry", func(t *testing.T) {
					user1_auth2, err = service.LoginInternal(ctx, user1_input.Email, user1_input.Password)
					if err != nil {
						t.Fatal("could not login", err.Error())
					}
					if user1_auth2.User.AuthStateID == user1_auth.User.AuthStateID {
						t.Fatal("auth_state.id should not match")
					}

					var logins []int
					qb := table.AuthState.SELECT(table.AuthState.ID).
						FROM(table.AuthState).
						WHERE(table.AuthState.UserID.EQ(postgres.Int(user1.ID)))
					if err := qb.QueryContext(ctx, db, &logins); err != nil {
						t.Fatal("could not query auth state for user", err.Error())
					}
					if len(logins) != 2 {
						t.Fatal("auth_state for user should be 2")
					}
				})
			})
		})

		t.Run("incorrect input", func(t *testing.T) {
			_, err := service.LoginInternal(ctx, user1_input.Email, "somerandompassword")
			if err == nil {
				t.Fatal("login should fail. password is incorrect")
			}
		})
	})
}
