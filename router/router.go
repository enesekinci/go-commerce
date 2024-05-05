package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-commerce/handler"
	"go-commerce/middleware"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	api.Get("/", handler.Hello)

	authRoutes(api)

	userRoutes(api)

	coreRoutes(api)
}

func authRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/token/refresh", handler.RefreshToken)
	auth.Post("/token/access", middleware.VerifyRefreshToken(), handler.AccessToken)
	auth.Post("/password/forgot", handler.ForgotPassword)
	auth.Post("/password/change", middleware.VerifyAccessToken(), handler.ChangePassword)
	auth.Post("/profile/update", middleware.VerifyAccessToken(), handler.UpdateProfile)
}

func userRoutes(api fiber.Router) {
	user := api.Group("/user", middleware.VerifyAccessToken())
	user.Get("/:id", handler.GetUser)
	user.Post("/create", handler.CreateUser)
	user.Post("/update/:id", middleware.VerifyAccessToken(), middleware.IsAdmin, handler.UpdateUser)
	user.Get("/delete/:id", middleware.VerifyAccessToken(), handler.DeleteUserYourself)

	user.Get("/delete/admin/:id", middleware.VerifyAccessToken(), middleware.IsAdmin, handler.DeleteUserByAdmin)
}

func coreRoutes(api fiber.Router) {
	api.Get("/error-message", handler.ErrorMessage)
	api.Get("/constant", handler.Constant)
	api.Get("/province", handler.Province)
	api.Get("/city/:province_id", handler.City)
	api.Get("/district/:city_id", handler.District)
	api.Get("/currency", handler.Currency)
}
