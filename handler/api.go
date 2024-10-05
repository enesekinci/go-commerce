package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// Hello handle api status
func Hello(c *fiber.Ctx) error {

	fmt.Println("Hello")

	return c.JSON(fiber.Map{"status": true, "message": "Hello i'm ok!", "data": nil})
}
