package main

import (
	"blogger-api/config"
	"blogger-api/controllers"
	"blogger-api/repositories"
	"blogger-api/services"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()

	blogRepo := repositories.NewBlogRepository(config.DB)
	blogService := services.NewBlogService(blogRepo)
	blogController := controllers.NewBlogController(blogService)

	app := fiber.New()

	api := app.Group("/api")
	api.Post("/blogs", blogController.CreateBlog)
	api.Get("/blogs", blogController.GetAll)
	api.Get("/blogs/:id", blogController.GetById)
	api.Put("/blogs/:id", blogController.Update)
	api.Delete("/blogs/:id", blogController.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}