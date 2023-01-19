package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Convert(c *fiber.Ctx) error {

	c.JSON(struct {
		Amount float64 `json:"amount"`
		Asset  string  `json:"asset"`
	}{
		Amount: 5.38,
		Asset:  "BRL",
	})
	return nil
}
