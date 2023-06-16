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
	// Arrange
	app := fiber.New()
	app.Get("/protected", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("Protected route")
	})

	// Set up a valid token for testing
	token := generateValidToken()

	// Create a request with the token in the cookie
	req := httptest.NewRequest(fiber.MethodGet, "/protected", nil)
	req.Header.Set("Cookie", "jwt="+token)

	// Create a response recorder to capture the response
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Protected route", string(body))
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	// Arrange
	app := fiber.New()
	app.Get("/protected", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("Protected route")
	})

	// Set up an invalid token for testing
	token := "invalid-token"

	// Create a request with the token in the cookie
	req := httptest.NewRequest(fiber.MethodGet, "/protected", nil)
	req.Header.Set("Cookie", "jwt="+token)

	// Create a response recorder to capture the response
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, `{"message":"unauthorized"}`, string(body))
}

func generateValidToken() string {
	claims := jwt.StandardClaims{
		// Set the claims as per your requirements
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Set the SecretKey as per your application's secret key
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return tokenString
}
