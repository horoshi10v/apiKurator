package test

import (
	"apiKurator/database"
	"apiKurator/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	os.Setenv("DATABASE_CONFIG", "root:root@/eKurator?&parseTime=True")

	database.Connect()

	assert.NotNil(t, database.DB)

	err := database.DB.AutoMigrate(&models.User{})
	assert.NoError(t, err)

	sqlDB, err := database.DB.DB()
	assert.NoError(t, err)
	sqlDB.Close()
}
