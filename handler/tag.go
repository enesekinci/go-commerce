package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-commerce/app/models"
	"go-commerce/database"
)

type CreateTagInput struct {
	Name string `validate:"required,min=3,max=100" json:"name"`
}

func GetTags(c *fiber.Ctx) error {
	var tags []models.Tag

	database.DB.Find(&tags)

	return c.JSON(fiber.Map{"status": true, "message": "Tags found", "tags": tags})
}

func CreateTag(c *fiber.Ctx) error {

	input := new(CreateTagInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	if database.IsExistInDB("tags", "name", input.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Tag already exist", "data": nil})
	}

	tag := models.NewTag(input.Name)

	database.DB.Create(&tag)

	return c.JSON(fiber.Map{"status": true, "message": "Tag created", "tag": tag})
}

func DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")

	var tag models.Tag

	database.DB.First(&tag, id)

	if tag.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "No tag found with ID", "data": nil})
	}

	database.DB.Delete(&tag)

	return c.JSON(fiber.Map{"status": true, "message": "Tag deleted", "tag": tag})
}
