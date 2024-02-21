package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ayaanqui/go-migration-tool/migration_tool"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/stadio-app/stadio-backend/database"
	"github.com/stadio-app/stadio-backend/graph"
	gresolver "github.com/stadio-app/stadio-backend/graph/resolver"
	"github.com/stadio-app/stadio-backend/services"
	"github.com/stadio-app/stadio-backend/types"
	"github.com/stadio-app/stadio-backend/utils"
)

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

	server.Tokens = tokens
	service := services.Service{
		DB: server.DB,
		StructValidator: server.StructValidator,
		Tokens: tokens,
	}

	// TODO: only show in dev environment
	server.Router.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	server.Router.Group(func(r chi.Router) {
		c := graph.Config{
			Resolvers: &gresolver.Resolver{
				AppContext: server,
				Service: service,
			},
		}
		graphql_handler := handler.NewDefaultServer(graph.NewExecutableSchema(c))
		r.Handle("/graphql", graphql_handler)
	})
	return &server
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := database.NewDbConnection()
	if err != nil {
		log.Println("❌ Could not connect to database.")
		panic(err)
	}
	defer db.Close()
	router := chi.NewRouter()
	server := NewServer(db, router)

	log.Printf("🚀 Server running http://localhost:%s/playground\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), server.Router); err != nil {
		log.Fatal("❌ Could not start server", err)
	}
}
