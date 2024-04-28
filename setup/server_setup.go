package setup

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ayaanqui/go-migration-tool/migration_tool"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/sendgrid/sendgrid-go"
	"github.com/stadio-app/stadio-backend/graph"
	gresolver "github.com/stadio-app/stadio-backend/graph/resolver"
	"github.com/stadio-app/stadio-backend/services"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"
)

const GRAPH_ENDPOINT string = "/graphql"

func NewServer(db_conn *sql.DB, router *chi.Mux) *types.ServerBase {
	server := types.ServerBase{
		DB: db_conn,
		Router: router,
		StructValidator: validator.New(),
		MigrationDirectory: "./database/migrations",
	}

	// Run DB migrations
	db_migration := migration_tool.New(server.DB, &migration_tool.Config{
		Directory: server.MigrationDirectory,
		TableName: "migration",
	})
	db_migration.RunMigration()

	var tokens types.Tokens
	if err := utils.FileMapper("./tokens.json", &tokens); err != nil {
		panic(err)
	}

	server.Tokens = &tokens

	// Cloudinary CDN
	cloudinary, err := cloudinary.NewFromParams(tokens.Cloudinary.CloudName, tokens.Cloudinary.ApiKey, tokens.Cloudinary.ApiSecret)
	if err != nil {
		panic(err)
	}

	sendgrid_client := sendgrid.NewSendClient(tokens.SendGrid.ApiKey)

	service := services.Service{
		DB: server.DB,
		StructValidator: server.StructValidator,
		Tokens: &tokens,
		Cloudinary: cloudinary,
		Sendgrid: sendgrid_client,
	}

	// TODO: only show in dev environment
	server.Router.Handle("/playground", playground.Handler("GraphQL Playground", GRAPH_ENDPOINT))

	server.Router.Group(func(chi_router chi.Router) {
		c := graph.Config{}
		c.Resolvers = &gresolver.Resolver{
			AppContext: server,
			Service: service,
		}
		c.Directives.IsAuthenticated = service.IsAuthenticatedDirective
		graphql_handler := handler.NewDefaultServer(graph.NewExecutableSchema(c))

		chi_router.Use(service.AuthorizationMiddleware)
		chi_router.Handle(GRAPH_ENDPOINT, graphql_handler)
	})
	return &server
}
