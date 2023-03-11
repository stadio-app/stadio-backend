package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/m3-app/backend/app"
	"github.com/m3-app/backend/utils"
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
	
	http.Handle("/graphql", app.GqlServer)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", app.Port)
	log.Fatal(http.ListenAndServe(port_str, nil))
}
