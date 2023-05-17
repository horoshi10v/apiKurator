package main

import (
	"apiKurator/config"
	"apiKurator/database"
	"apiKurator/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {

	database.Connect()

	app := fiber.New()

	config.GoogleConfig()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
