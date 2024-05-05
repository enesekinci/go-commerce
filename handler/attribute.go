package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-commerce/app/models"
	"go-commerce/database"
)

type CreateAttributeInput struct {
	Name       string `validate:"required,min=3,max=100" json:"name"`
	Filterable uint   `json:"filterable"`
	Values     string `json:"values"`
	Categories []uint `json:"categories"`
}

func GetAttributes(c *fiber.Ctx) error {
	var attributes []models.Attribute

	database.DB.Find(&attributes).Preload("Categories")

	return c.JSON(fiber.Map{"status": "success", "message": "Attributes found", "attributes": attributes})
}

func CreateAttribute(c *fiber.Ctx) error {
	input := new(CreateAttributeInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your request", "error": err})
	}

	attribute := models.NewAttribute(input.Name, input.Filterable, input.Values)

	database.DB.Create(&attribute)

	if len(input.Categories) > 0 {
		for _, categoryID := range input.Categories {
			var category models.Category

			database.DB.First(&category, categoryID)

			if category.Name != "" {
				err := database.DB.Model(&attribute).Association("Categories").Append(&category)

				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to append category", "error": err})
				}
			}
		}
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Attribute created", "attribute": attribute})
}

func UpdateAttribute(c *fiber.Ctx) error {
	input := new(CreateAttributeInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your request", "error": err})
	}

	id := c.Params("id")

	var attribute models.Attribute

	database.DB.First(&attribute, id)

	if attribute.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No attribute found with ID", "data": nil})
	}

	attribute.Name = input.Name
	attribute.Filterable = input.Filterable
	attribute.Values = []byte(input.Values)

	database.DB.Save(&attribute)

	if len(input.Categories) > 0 {
		err := database.DB.Model(&attribute).Association("Categories").Clear()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to clear categories", "error": err})
		}

		for _, categoryID := range input.Categories {
			var category models.Category

			database.DB.First(&category, categoryID)

			if category.Name != "" {
				err := database.DB.Model(&attribute).Association("Categories").Append(&category)

				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to append category", "error": err})
				}
			}
		}
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Attribute updated", "attribute": attribute})
}

func DeleteAttribute(c *fiber.Ctx) error {
	id := c.Params("id")

	var attribute models.Attribute

	database.DB.First(&attribute, id)

	if attribute.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No attribute found with ID", "data": nil})
	}

	database.DB.Delete(&attribute)

	return c.JSON(fiber.Map{"status": "success", "message": "Attribute deleted", "attribute": attribute})
}
