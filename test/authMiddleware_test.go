package test

import (
	"apiKurator/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	app := fiber.New()
	app.Get("/protected", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("Protected route")
	})
	token := generateValidToken()

	req := httptest.NewRequest(fiber.MethodGet, "/protected", nil)
	req.Header.Set("Cookie", "jwt="+token)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Protected route", string(body))
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {

	app := fiber.New()
	app.Get("/protected", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("Protected route")
	})
	token := "invalid-token"

	req := httptest.NewRequest(fiber.MethodGet, "/protected", nil)
	req.Header.Set("Cookie", "jwt="+token)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, `{"message":"unauthorized"}`, string(body))
}

func generateValidToken() string {
	claims := jwt.StandardClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return tokenString
}
