package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Response struct {
	Message string `json:"message"`
}

const PORT = 3000

func main() {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Chicago", 
		"localhost",
		"postgres",
		"password",
		"postgres",
		"5432",
		"verify-full",
	)
	_, err := sql.Open("postgres", dns)
	if err != nil {
		panic(err)
	}
	
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