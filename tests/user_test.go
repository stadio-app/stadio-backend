package tests

import (
	"context"
	"testing"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func TestUser(t *testing.T) {
	app, service := NewMockServer(t)
	var user1 gmodel.User
	
	t.Run("create user", func(t *testing.T) {
		var err error
		user1, err = service.CreateInternalUser(context.Background(), gmodel.CreateAccountInput{
			Email: "user1@email.com",
			Name: "User 1",
			Password: "password123",
		})
		if err != nil {
			t.Fatal(err)
		}
		if *user1.Active {
			t.Fatal("User.Active should be set to false")
		}
		if *user1.AuthPlatform != gmodel.AuthPlatformTypeInternal {
			t.Fatal("user auth platform should be internal")
		}
	})

	t.Cleanup(func() {
		_, err := table.User.DELETE().WHERE(postgres.Bool(true)).Exec(app.DB)
		if err != nil {
			t.Fatal(err)
		}
	})
}
