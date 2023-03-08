package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/m3-app/backend/util"
)

type Response struct {
	Message string `json:"message"`
}

const PORT = 3000

func main() {
	db_conn := util.DbConnection()
	defer db_conn.Close()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Response{
			Message: "Hello",
		})
	})

	if err := app.Listen(fmt.Sprintf(":%d", PORT)); err != nil {
		log.Fatalf("port %d is already in use", PORT)
	}
}
