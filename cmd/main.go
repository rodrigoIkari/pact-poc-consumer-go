package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", index)

	app.Listen(":3000")
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
