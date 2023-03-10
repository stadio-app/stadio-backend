package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/m3-app/backend/app"
	"github.com/m3-app/backend/utils"
)

const PORT uint = 8080

func main() {
	db_conn := utils.DbConnection()
	defer db_conn.Close()

	server := fiber.New()
	app := app.New(db_conn, server)

	if err := app.Server.Listen(fmt.Sprintf(":%d", PORT)); err != nil {
		log.Fatalf("port %d is already in use", PORT)
	}
}
