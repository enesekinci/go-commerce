package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-commerce/app/models"
	"go-commerce/database"
)

type CreateOptionInput struct {
	Name   string `validate:"required,min=3,max=100" json:"name"`
	TypeId uint   `json:"type_id" validate:"required"`
	Values string `json:"values" validate:"required"`
}

func GetOptions(c *fiber.Ctx) error {
	var options []models.Option

	database.DB.Preload("Values").Find(&options)

	return c.JSON(fiber.Map{"status": "success", "message": "All options", "data": options})

}

func CreateOption(c *fiber.Ctx) error {
	var data CreateOptionInput

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	option := models.NewOption(data.Name, data.TypeId, data.Values)

	database.DB.Create(&option)

	return c.JSON(fiber.Map{"status": "success", "message": "Option created", "data": option})
}

func UpdateOption(c *fiber.Ctx) error {
	input := new(CreateOptionInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	id := c.Params("id")

	var option models.Option

	database.DB.First(&option, id)

	if option.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Option not found", "data": nil})
	}

	option.Name = input.Name
	option.TypeId = input.TypeId
	option.Values = []byte(input.Values)

	database.DB.Save(&option)

	return c.JSON(fiber.Map{"status": "success", "message": "Option updated", "data": option})
}

func DeleteOption(c *fiber.Ctx) error {
	id := c.Params("id")

	var option models.Option

	database.DB.First(&option, id)

	if option.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Option not found", "data": nil})
	}

	database.DB.Delete(&option)

	return c.JSON(fiber.Map{"status": "success", "message": "Option deleted", "data": nil})
}
