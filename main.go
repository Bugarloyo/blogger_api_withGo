package main

import (
	"blogger-api/config"
	"blogger-api/controllers"
	"blogger-api/repositories"
	"blogger-api/services"
	"blogger-api/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.ConnectDB()

	blogRepo := repositories.NewBlogRepository(config.DB)
	blogService := services.NewBlogService(blogRepo)
	blogController := controllers.NewBlogController(blogService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:    "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:   "*",
		AllowCredentials: false,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	}))

	routes.SetupBlogRoutes(app, blogController)

	

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}