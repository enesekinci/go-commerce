package helper

import "github.com/gofiber/fiber/v2"

func GetPaginateParams(c *fiber.Ctx) (string, string) {
	page := c.Query("page", "1")
	perPage := c.Query("perPage", "5")

	if StringToInt(perPage) > 100 || StringToInt(perPage) < 1 {
		perPage = "5"
	}

	if StringToInt(page) < 1 {
		page = "1"
	}
	return page, perPage
}

func GetSearchParams(c *fiber.Ctx) (string, string) {
	search := c.Query("search", "")
	field := c.Query("searchField", "")
	return search, field
}
