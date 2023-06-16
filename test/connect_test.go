package test

import (
	"apiKurator/database"
	"apiKurator/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// Set up your test environment
	// For example, you can use a separate test database or a test configuration

	// Set the DATABASE_CONFIG environment variable to the test database connection string
	os.Setenv("DATABASE_CONFIG", "root:root@/eKurator?&parseTime=True")

	// Call the Connect function to establish a connection
	database.Connect()

	// Assert that the DB variable is not nil
	assert.NotNil(t, database.DB)

	// Migrate the User model to the database
	err := database.DB.AutoMigrate(&models.User{})
	assert.NoError(t, err)

	// Optional: Perform additional assertions or checks to ensure the database connection and migration were successful
	// For example, you can check if the User table exists or insert test data and retrieve it

	// Cleanup: Close the database connection and perform any necessary cleanup steps
	sqlDB, err := database.DB.DB()
	assert.NoError(t, err)
	sqlDB.Close()
}
