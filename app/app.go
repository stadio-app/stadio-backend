package app

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/m3-app/backend/controllers"
	"github.com/m3-app/backend/services"
	"github.com/m3-app/backend/types"
	"github.com/m3-app/backend/utils"
)

type AppBase struct {
	types.AppContext
	Services services.Service
}

func New(db_conn *sql.DB, server *fiber.App) *AppBase {
	app := AppBase{}
	app.DbConn = db_conn
	app.Server = server
	// initialize tokens
	tokens, tokens_err := utils.ParseTokens()
	if tokens_err != nil {
		panic(tokens_err)
	}
	app.Tokens = &tokens
	return app.NewBaseHandler()
}

func (app *AppBase) NewBaseHandler() *AppBase {
	app.EntityManager = utils.CreateEntClient(app.DbConn) // chain .Debug() to show messages
	// auto migration
	ctx := context.Background()
	if err := app.EntityManager.Schema.Create(ctx); err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	app.Validator = validator.New()
	app.Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	app.Services = services.New(services.ServiceConfig{
		DbConn: app.DbConn,
		EntityManager: app.EntityManager,
		Validator: app.Validator,
	})
	// utils.SetupOauthProviders(*app.Tokens) // oauth providers
	controllers.New(app.AppContext, app.Services)
	return app
}
