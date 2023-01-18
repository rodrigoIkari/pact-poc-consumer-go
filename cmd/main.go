package main

import (
	"github.com/gofiber/fiber/v2"
	r "github.com/rodrigoikari/pact-poc-consumer-go/api"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	r.ConfigRoutes(api)

	app.Listen(":3000")

}
