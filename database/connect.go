package database

import (
	"apiKurator/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {

	var configConn = os.Getenv("DATABASE_CONFIG")
	conn, err := gorm.Open(mysql.Open(configConn), &gorm.Config{})
	if err != nil {
		log.Panicln("can't open database connection")
	}

	DB = conn
	conn.AutoMigrate(&models.User{})
}
