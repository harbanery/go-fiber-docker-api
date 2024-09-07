package main

import (
	"go-fiber-docker-api/src/config"
	"go-fiber-docker-api/src/controllers"
	"go-fiber-docker-api/src/helpers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return "0.0.0.0:" + port

}

func main() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app := fiber.New()
	config.InitDB()
	helpers.Migration()
	ProductRoute(app)

	if err := app.Listen(getPort()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func ProductRoute(app *fiber.App) {
	products := app.Group("/products")
	products.Get("/", controllers.GetAllProducts)
	products.Get("/:id", controllers.GetDetailProduct)
	products.Post("/", controllers.AddProduct)
	products.Put("/:id", controllers.EditProduct)
	products.Delete("/:id", controllers.DeleteProduct)
}
