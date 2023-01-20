package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodrigoikari/pact-poc-consumer-go/controllers"
)

func ConfigRoutes(api fiber.Router) {
	exchangeApi := api.Group("/checkout/")
	exchangeApi.Post("simulation", controllers.SimulateCart).Name("Checkout Simulation")
}
