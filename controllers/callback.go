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
	"time"
)

func GoogleCallback(c *fiber.Ctx) error {
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
	// Set a cookie
	//c.Cookie(&fiber.Cookie{
	//	Name:     "MyCookie",
	//	Value:    fmt.Sprintf("%d", time.Now().Unix()),
	//	HTTPOnly: true,
	//	Expires:  time.Now().Add(time.Hour * 1),
	//})
	//database.DB.Create(&user)
	//
	//return c.JSON(user)
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

	return c.JSON(fiber.Map{
		"massage": "success",
	})
}
