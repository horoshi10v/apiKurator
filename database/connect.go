package database

import (
	"apiKurator/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:root@/eKurator?&parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Panicln("can't open database connection")
	}

	DB = conn
	conn.AutoMigrate(&models.User{})
}
