package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-commerce/app/models"
	"go-commerce/core/helper"
	"go-commerce/database"
)

//GetVariantTypes
//GetVariants
//CreateVariant
//UpdateVariant
//DeleteVariant

type CreateVariantInput struct {
	Name          string `validate:"required,min=3,max=100" json:"name"`
	VariantTypeID uint   `json:"variant_type_id"`
	Values        string `json:"values"`
}

func GetVariantTypes(c *fiber.Ctx) error {
	var variantTypes []models.VariantType

	database.DB.Find(&variantTypes)

	return c.JSON(fiber.Map{"status": true, "message": "Variant types found", "variant_types": variantTypes})
}

func GetVariants(c *fiber.Ctx) error {
	var variants []models.Variant

	database.DB.Preload("VariantType").Find(&variants)

	return c.JSON(fiber.Map{"status": true, "message": "Variants found", "variants": variants})
}

/*func GetVariant(c *fiber.Ctx) error {

	id := c.Params("id")

	var variant models.Variant

	database.DB.Find(&variant, id)

	if variant.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "No variant found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "Variant found", "variant": variant})
}*/

func CreateVariant(c *fiber.Ctx) error {
	input := new(CreateVariantInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	var variant models.Variant
	variant.Name = input.Name
	variant.VariantTypeID = input.VariantTypeID
	variant.Values = []byte(input.Values)

	database.DB.Create(&variant)

	return c.JSON(fiber.Map{"status": true, "message": "Variant created", "variant": variant})
}

func UpdateVariant(c *fiber.Ctx) error {
	id := c.Params("id")

	input := new(CreateVariantInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	var variant models.Variant

	database.DB.Find(&variant, id)

	if variant.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "No variant found with ID", "data": nil})
	}

	variant.Name = input.Name
	variant.VariantTypeID = input.VariantTypeID
	variant.Values = []byte(input.Values)

	database.DB.Save(&variant)

	return c.JSON(fiber.Map{"status": true, "message": "Variant updated", "variant": variant})
}

func DeleteVariant(c *fiber.Ctx) error {
	id := c.Params("id")

	var variant models.Variant

	database.DB.Find(&variant, id)

	if variant.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "No variant found with ID", "data": nil})
	}

	database.DB.Delete(&variant)

	return c.JSON(fiber.Map{"status": true, "message": "Variant deleted", "variant": variant})
}
