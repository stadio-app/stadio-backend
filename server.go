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
	"github.com/m3-app/backend/app"
	"github.com/m3-app/backend/graph"
	"github.com/m3-app/backend/utils"
	"github.com/markbates/goth/gothic"
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
	router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	// oauth routes
	router.Group(func(r chi.Router) {
		r.Use(app.GothMiddleware())
		r.HandleFunc("/auth/{provider:[a-z-]+}", func(w http.ResponseWriter, r *http.Request) {
			gothic.BeginAuthHandler(w, r)
		})
		r.HandleFunc("/auth/{provider:[a-z-]+}/callback", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			provider_user, err := gothic.CompleteUserAuth(w, r)
			if err != nil {
				utils.ErrorResponse(w, 400, "could not complete oauth transaction")
				return
			}
			utils.DataResponse(w, provider_user)
		})
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

	log.Printf("Server running on http://localhost:%s/\n", app.Port)
	if err := http.ListenAndServe(port_str, router); err != nil {
		log.Fatal(err)
	}
}
