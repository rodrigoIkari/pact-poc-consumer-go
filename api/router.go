package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodrigoikari/pact-poc-consumer-go/controllers"
)

func ConfigRoutes(api fiber.Router) {
	exchangeApi := api.Group("/exchange")
	exchangeApi.Post("convert", controllers.Convert).Name("Convert Rate")
}
