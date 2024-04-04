package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoikari/pact-poc-consumer-go/models"
	"github.com/rodrigoikari/pact-poc-consumer-go/services"
)

type CartController struct {
	currencyService services.CurrencyService
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
	totalAmount, err := ctl.CalculateCartValue(cart)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	fmt.Println("Cart Value: ", totalAmount, cart.Buyer_Asset_Code)

	cartResponse := new(models.CartResponse)
	cartResponse.Items = cart.Items
	cartResponse.Buyer_Asset_Code = cart.Buyer_Asset_Code
	cartResponse.Total_Amount = totalAmount

	return c.JSON(cartResponse)
}

func (ctl *CartController) CalculateCartValue(cart *models.CartRequest) (float64, error) {

	totalAmount := 0.0

	for _, it := range cart.Items {
		convertedValue, err := ctl.currencyService.ConvertCurrencyValue(it.Amount*float64(it.Quantity), it.Seller_Asset_Code, cart.Buyer_Asset_Code)
		if err == nil {
			return 0, err
		}
		totalAmount += convertedValue
	}

	return totalAmount, nil
}
