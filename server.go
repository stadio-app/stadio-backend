package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/m3-app/backend/util"
)

type Response struct {
	Message string `json:"message"`
}

const PORT uint = 3000

func main() {
	db_conn := util.DbConnection()
	defer db_conn.Close()

	ent_manager := util.CreateEntClient(db_conn).Debug()
	ent_manager.Location.Create()

	server := fiber.New()
	server.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Response{
			Message: "Hello",
		})
	})

	ctx := context.Background()

	if err := ent_manager.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := server.Listen(fmt.Sprintf(":%d", PORT)); err != nil {
		log.Fatalf("port %d is already in use", PORT)
	}
}
