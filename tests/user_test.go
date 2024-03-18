package tests

import (
	"testing"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func TestUser(t *testing.T) {
	var user1 gmodel.User
	user1_input := gmodel.CreateAccountInput{
		Email: "user1@email.com",
		Name: "User 1",
		Password: "password123",
	}
	
	t.Run("create user", func(t *testing.T) {
		var err error
		user1, err = service.CreateInternalUser(ctx, user1_input)
		if err != nil {
			t.Fatal(err)
		}
		if *user1.Active {
			t.Fatal("User.Active should be set to false")
		}
		if *user1.AuthPlatform != gmodel.AuthPlatformTypeInternal {
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

}
