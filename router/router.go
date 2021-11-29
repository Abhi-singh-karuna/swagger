package router

import (
	"path/handler"

	"github.com/gofiber/fiber/v2"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/docs/*", swagger.Handler)
	api := app.Group("/product")

	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)

}
