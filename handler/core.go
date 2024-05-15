package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-commerce/core/constant"
	"strconv"
)

func ErrorMessage(c *fiber.Ctx) error {

	//var messages map[int]string

	//for _, message := range constant.AllErrorCodes() {
	//	code := int(message)
	//	messages[code] = message.String()
	//}

	errorCodes := constant.AllErrorCodes()

	errors := make(map[string]string)

	for _, code := range errorCodes {
		errors[strconv.Itoa(int(code))] = code.String()
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":        "success",
		"message":       "All error messages",
		"errorMessages": errors,
	})
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
