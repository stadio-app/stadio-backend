package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
}

const PORT = 3000

func main() {
	app := fiber.New()
	app.Get("/", func (c *fiber.Ctx) error {
        return c.JSON(Response{
			Message: "Hello",
		})
    })

	if err := app.Listen(fmt.Sprintf(":%d", PORT)); err != nil {
		log.Fatalf("Could not start server. Port %d is already in use", PORT)
	}
}