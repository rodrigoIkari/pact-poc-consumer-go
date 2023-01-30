package controllers

import (
	"fmt"
	"strings"

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
		return c.SendStatus(400)
	}

	if len(cart.Items) <= 0 {
		fmt.Println("error processing cart: No Items")
		return c.SendStatus(422)
	}

	if strings.TrimSpace(cart.Buyer_Asset_Code) == "" {
		fmt.Println("error processing cart: No Buyer Asset Code")
		return c.SendStatus(422)
	}

	fmt.Println("cart request accepted")

	fmt.Println("Calculating Cart Value ...")

	ctl.CalculateCartValue()

	return c.JSON(cart)
}

func (ctl *CartController) CalculateCartValue() (float64, error) {
	return 0, nil
}
