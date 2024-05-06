package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-commerce/database"
	"go-commerce/router"
	"log"
)

func main() {
	//TODO: MODELLERIN ILISKILERI DUZENLENECEK, CASCADE VEYA SET NULL DURUMLARI AYARLANACAK

	app := fiber.New(fiber.Config{
		AppName:     "Go Commerce",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
