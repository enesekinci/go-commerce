package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorMessage(c *fiber.Ctx) error {
	return c.SendString("Error Messages")
}
func Constant(c *fiber.Ctx) error {
	return c.SendString("Constants")
}
func Province(c *fiber.Ctx) error {
	return c.SendString("Provinces")
}
func City(c *fiber.Ctx) error {
	return c.SendString("Cities")
}
func District(c *fiber.Ctx) error {
	return c.SendString("Districts")
}
func Currency(c *fiber.Ctx) error {
	return c.SendString("Currencies")
}
