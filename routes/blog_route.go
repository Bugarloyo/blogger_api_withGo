package routes

import (
	"blogger-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupBlogRoutes(app *fiber.App, blogController *controllers.BlogController) {
	api := app.Group("/api")

	api.Post("/blogs", blogController.CreateBlog)
	api.Get("/blogs", blogController.GetAll)
	api.Get("/blogs/:id", blogController.GetById)
	api.Put("/blogs/:id", blogController.Update)
	api.Delete("/blogs/:id", blogController.Delete)
}