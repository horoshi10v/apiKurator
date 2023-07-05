package controllers

import (
	"apiKurator/config"
	"apiKurator/database"
	"apiKurator/models"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"time"
)

func (u *UserControllerImpl) GoogleLogin(c *fiber.Ctx) error {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	c.Status(fiber.StatusSeeOther)
	err := c.Redirect(url)
	if err != nil {
		return err
	}
	return c.JSON(url)
}

func (u *UserControllerImpl) GoogleLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logout success",
	})
}

func (u *UserControllerImpl) GoogleCallback(c *fiber.Ctx) error {
	var clientPort = os.Getenv("CLIENT_PORT")
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	googlecon := config.GoogleConfig()

	gtoken, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + gtoken.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	var user models.User
	err = json.NewDecoder(resp.Body).Decode(&user)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    string(user.ID),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"massage": "can not sign",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	database.DB.Create(&user)

	return c.Redirect(clientPort + "/user")
}
