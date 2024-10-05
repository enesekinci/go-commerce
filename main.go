package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"go-commerce/database"
	"go-commerce/router"
	"log"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	//TODO: MODELLERIN ILISKILERI DUZENLENECEK, CASCADE VEYA SET NULL DURUMLARI AYARLANACAK

	app := fiber.New(fiber.Config{
		AppName:     "Go Commerce",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
