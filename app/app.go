package app

import (
	"database/sql"
	"reflect"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"github.com/stadio-app/stadio-backend/ent"
	"github.com/stadio-app/stadio-backend/services"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"
)

type AppBase struct {
	types.AppContext
	Services services.Service
}

func New(db_conn *sql.DB, port string) *AppBase {
	app := AppBase{}
	app.Port = port
	
	app.DbConn = db_conn
	app.EntityDriver = entsql.OpenDB(dialect.Postgres, app.DbConn)
	app.EntityManager = ent.NewClient(ent.Driver(app.EntityDriver))

	// initialize tokens
	tokens, tokens_err := utils.ParseTokens()
	if tokens_err != nil {
		panic(tokens_err)
	}
	app.Tokens = &tokens
	return app.NewBaseHandler()
}

func (app *AppBase) NewBaseHandler() *AppBase {
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
	utils.SetupOauthProviders(*app.Tokens)
	return app
}
