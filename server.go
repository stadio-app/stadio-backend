package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/m3-app/backend/app"
	"github.com/m3-app/backend/graph"
	"github.com/m3-app/backend/utils"
)

const defaultPort = "8080"

func main() {
	// GraphQL integration https://entgo.io/docs/graphql#quick-introduction
	entc_generate()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	
	db_conn := utils.DbConnection()
	defer db_conn.Close()
	
	app := app.New(db_conn, port)
	port_str := fmt.Sprintf(":%s", app.Port)

	router := chi.NewRouter()
	router.Use(app.BaseMiddleware())
	router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	// secure routes
	router.Group(func(r chi.Router) {
		gql_server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{
				AppBase: app,
			},
		}))
		r.Use(app.AuthMiddleware())
		r.Handle("/graphql", gql_server)
	})

	log.Printf("Server running on http://localhost:%s/\n", app.Port)
	if err := http.ListenAndServe(port_str, router); err != nil {
		log.Fatal(err)
	}
}

func entc_generate() {
	ex, err := entgql.NewExtension()
    if err != nil {
        log.Fatalf("creating entgql extension: %v", err)
    }
    if err := entc.Generate("./ent/schema", &gen.Config{}, entc.Extensions(ex)); err != nil {
        log.Fatalf("running ent codegen: %v", err)
    }
}
