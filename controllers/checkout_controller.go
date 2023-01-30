package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoikari/pact-poc-consumer-go/models"
	"github.com/rodrigoikari/pact-poc-consumer-go/services"
)

type CartController struct {
	currencyService *services.CurrencyService
}

func (ctl *CartController) SimulateCart(c *fiber.Ctx) error {

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

	fmt.Println("Calculating Cart Value ...")

	ctl.CalculateCartValue()

	return c.JSON(cart)
}

func (ctl *CartController) CalculateCartValue() (float64, error) {
	return 0, nil
}
