package routes

import (
	"apiKurator/controllers"
	"apiKurator/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	userController := &controllers.UserControllerImpl{}

	app.Get("/google_login", userController.GoogleLogin)
	app.Get("/google_callback", userController.GoogleCallback)
	app.Get("/google_logout", userController.GoogleLogout)

	app.Get("/user", middleware.AuthMiddleware, userController.User)
	app.Get("/users", userController.GetUsers)
	app.Get("user/:id", userController.GetUser)
	app.Put("user/:id/update", userController.UpdateUser)
	app.Delete("/user/:id", userController.DeleteUser)
	app.Post("/addUser", userController.AddUser)
}
