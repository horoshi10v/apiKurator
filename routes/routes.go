package routes

import (
	"apiKurator/controllers"
	"apiKurator/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/google_login", controllers.GoogleLogin)
	app.Get("/google_callback", controllers.GoogleCallback)
	app.Get("/g_logout", controllers.GoogleLogout)

	app.Get("/user", middleware.AuthMiddleware, controllers.User)
	app.Get("/users", controllers.GetUsers)
	app.Get("user/:id", controllers.GetUser)
	app.Put("user/:id/update", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)
}
