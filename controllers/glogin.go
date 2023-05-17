package controllers

import (
	"apiKurator/config"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GoogleLogin(c *fiber.Ctx) error {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func GoogleLogout(c *fiber.Ctx) error {
	//c.Cookie(&fiber.Cookie{
	//	Name:     "MyCookie",
	//	HTTPOnly: true,
	//	Expires:  time.Now().Add(-(time.Hour * 1000)),
	//})
	//// Redirect the user to the Google logout URL
	//return c.Redirect("https://accounts.google.com/logout")
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"massage": "logout success",
	})
}
