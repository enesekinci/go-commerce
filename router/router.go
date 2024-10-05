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

	categoryRoutes(api)

	brandRoutes(api)

	tagRoutes(api)

	attributeRoutes(api)

	variantRoutes(api)

	optionRoutes(api)

	productRoutes(api)
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

func categoryRoutes(api fiber.Router) {
	category := api.Group("/category")
	category.Get("/", handler.GetCategories)
	category.Get("/:id", handler.GetCategory)
	category.Post("/create", middleware.VerifyAccessToken(), handler.CreateCategory)
	category.Post("/update/:id", middleware.VerifyAccessToken(), handler.UpdateCategory)
	category.Get("/delete/:id", middleware.VerifyAccessToken(), handler.DeleteCategory)
}

func brandRoutes(api fiber.Router) {
	brand := api.Group("/brands")
	brand.Get("/", handler.GetBrands)
	brand.Get("/:id", handler.GetBrand)
	brand.Post("/create", middleware.VerifyAccessToken(), handler.CreateBrand)
	brand.Post("/update/:id", middleware.VerifyAccessToken(), handler.UpdateBrand)
	brand.Get("/delete/:id", middleware.VerifyAccessToken(), handler.DeleteBrand)
}

func tagRoutes(api fiber.Router) {
	tag := api.Group("/tag")
	tag.Get("/", handler.GetTags)
	tag.Post("/create", middleware.VerifyAccessToken(), handler.CreateTag)
	tag.Get("/delete/:id", middleware.VerifyAccessToken(), handler.DeleteTag)
}

func attributeRoutes(api fiber.Router) {
	attribute := api.Group("/attribute")
	attribute.Get("/", handler.GetAttributes)
	attribute.Post("/create", middleware.VerifyAccessToken(), handler.CreateAttribute)
	attribute.Post("/update/:id", middleware.VerifyAccessToken(), handler.UpdateAttribute)
	attribute.Get("/delete/:id", middleware.VerifyAccessToken(), handler.DeleteAttribute)
}

func variantRoutes(api fiber.Router) {
	variant := api.Group("/variant")

	variant.Get("/type", handler.GetVariantTypes)

	variant.Get("/", handler.GetVariants)
	variant.Post("/create", middleware.VerifyAccessToken(), handler.CreateVariant)
	variant.Post("/update/:id", middleware.VerifyAccessToken(), handler.UpdateVariant)
	variant.Get("/delete/:id", middleware.VerifyAccessToken(), handler.DeleteVariant)

}

func optionRoutes(api fiber.Router) {
	option := api.Group("/option")
	option.Get("/", handler.GetOptions)
	option.Post("/create", middleware.VerifyAccessToken(), handler.CreateOption)
	option.Post("/update/:id", middleware.VerifyAccessToken(), handler.UpdateOption)
	option.Get("/delete/:id", middleware.VerifyAccessToken(), handler.DeleteOption)
}

func productRoutes(api fiber.Router) {
	product := api.Group("/product")
	product.Get("/", handler.GetProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/create", middleware.VerifyAccessToken(), handler.CreateProduct)
	product.Post("/update/:id", middleware.VerifyAccessToken(), handler.UpdateProduct)
	product.Get("/delete/:id", middleware.VerifyAccessToken(), handler.DeleteProduct)
}
