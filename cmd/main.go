package main

import (
	"apiKurator/config"
	"apiKurator/controllers"
	"apiKurator/database"
	"apiKurator/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
)

func main() {

	config.GoogleConfig()

	var clientPort = os.Getenv("CLIENT_PORT")
	var serverPort = os.Getenv("SERVER_PORT")

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     clientPort,
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type",
		AllowCredentials: true,
	}))

	userController := &controllers.UserControllerImpl{}
	routes.Setup(app, userController)

	err := app.Listen(serverPort)
	if err != nil {
		log.Fatal(err)
	}
}
