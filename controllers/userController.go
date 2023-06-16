package controllers

import (
	"apiKurator/database"
	"apiKurator/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func AddUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	UserId := uuid.New().String()
	var user = models.User{
		ID:          UserId,
		Name:        data["name"],
		Email:       data["email"],
		Picture:     data["picture"],
		Role:        data["role"],
		Stage:       data["stage"],
		Department:  data["department"],
		Interests:   data["interests"],
		Description: data["description"],
		Phone:       data["phone"],
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	role := c.Query("role")

	var users []models.User
	//users?role=
	switch role {
	case "curator":
		database.DB.Where("role = ?", "куратор").Find(&users)
	case "student":
		database.DB.Where("role = ?", "студент").Find(&users)
	default:
		database.DB.Find(&users)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User
	if err := database.DB.Find(&user, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User
	if err := database.DB.Find(&user, "id = ?", userID).Error; err != nil {
		// Handle the error, e.g., user not found
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if val, ok := data["stage"].(string); ok {
		user.Stage = val
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
	if val, ok := data["phone"].(string); ok {
		user.Phone = val
	}
	if val, ok := data["name"].(string); ok {
		user.Name = val
	}
	if val, ok := data["picture"].(string); ok {
		user.Picture = val
	}
	if val, ok := data["email"].(string); ok {
		user.Email = val
	}

	database.DB.Save(&user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
