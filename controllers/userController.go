package controllers

import (
	"apiKurator/database"
	"apiKurator/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
)

var SecretKey = os.Getenv("SECRET_KEY")

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User

	database.DB.Where("id=?", claims.Issuer).First(&user)
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	role := c.Query("role") // Retrieve the role from the query parameter

	var users []models.User

	switch role {
	case "student":
		database.DB.Where("role = ?", "student").Find(&users)
	case "teacher":
		database.DB.Where("role = ?", "teacher").Find(&users)
	case "admin":
		database.DB.Where("role = ?", "admin").Find(&users)
	default:
		database.DB.Find(&users)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := database.DB.First(&user, "id = ?", id)
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		// Handle the error, e.g., user not found
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if val, ok := data["role"].(string); ok {
		user.Role = val
	}
	if val, ok := data["department"].(string); ok {
		user.Department = val
	}
	if val, ok := data["interests"].(string); ok {
		user.Interests = val
	}
	if val, ok := data["description"].(string); ok {
		user.Description = val
	}
	database.DB.Save(&user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	// Check if the user exists
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Delete the user
	database.DB.Delete(&user)

	return c.JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}
