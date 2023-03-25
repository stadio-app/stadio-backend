package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/stadio-app/stadio-backend/app"
	"github.com/stadio-app/stadio-backend/graph"
	"github.com/stadio-app/stadio-backend/types"
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
		c := graph.Config{
			Resolvers: &graph.Resolver{
				AppBase: app,
			},
		}
		c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
			var err error
			auth_header := ctx.Value(types.AuthHeader).(string)
			ctx, err = app.BearerAuthentication(ctx, auth_header)
			if err != nil {
				return nil, fmt.Errorf("unauthorized")
			}
			return next(ctx)
		}
		gql_server := handler.NewDefaultServer(graph.NewExecutableSchema(c))
		r.Use(app.BaseMiddleware())
		r.Handle("/graphql", gql_server)
	})

	log.Printf("🚀 Server running http://localhost:%s/playground\n", app.Port)
	if err := http.ListenAndServe(port_str, router); err != nil {
		log.Fatal(err)
	}
}
