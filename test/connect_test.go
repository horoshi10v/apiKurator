package test

import (
	"apiKurator/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	database.Connect()

	// Assert that the DB variable is not nil
	assert.NotNil(t, database.DB)
}
