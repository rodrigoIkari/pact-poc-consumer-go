package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoikari/pact-poc-consumer-go/models"
	"github.com/rodrigoikari/pact-poc-consumer-go/services"
)

func SimulateCart(c *fiber.Ctx) error {

	cart := new(models.CartRequest)

	if err := c.BodyParser(cart); err != nil {
		fmt.Println("error parsing cart: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Checkout Simulation: Bad Request")
	}

	errors := services.Validate(cart)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	fmt.Println("cart request accepted")
	return c.JSON(cart)
}
