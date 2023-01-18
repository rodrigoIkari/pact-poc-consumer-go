package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Convert(c *fiber.Ctx) error {

	c.JSON(struct {
		Amount float64
		Asset  string
	}{
		Amount: 5.38,
		Asset:  "BRL",
	})
	return nil
}
