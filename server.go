package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/stadio-app/stadio-backend/app"
	"github.com/stadio-app/stadio-backend/graph"
	"github.com/stadio-app/stadio-backend/utils"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	
	db_conn := utils.DbConnection()
	defer db_conn.Close()
	
	app := app.New(db_conn, port)
	port_str := fmt.Sprintf(":%s", app.Port)

	router := chi.NewRouter()
	router.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))
	// oauth routes
	router.Group(func(r chi.Router) {
		r.Use(app.GothMiddleware())
		r.HandleFunc("/auth/{provider}", app.OAuthSignIn)
		r.HandleFunc("/auth/{provider}/callback", app.OAuthCallback)
	})
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

	log.Printf("🚀 Server running http://localhost:%s/playground\n", app.Port)
	if err := http.ListenAndServe(port_str, router); err != nil {
		log.Fatal(err)
	}
}
