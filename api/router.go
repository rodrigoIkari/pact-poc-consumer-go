package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodrigoikari/pact-poc-consumer-go/controllers"
)

func ConfigRoutes(api fiber.Router) {
	exchangeApi := api.Group("/checkout/")

	ctl := new(controllers.CartController)

	exchangeApi.Post("simulation", ctl.SimulateCart).Name("Checkout Simulation")
}
