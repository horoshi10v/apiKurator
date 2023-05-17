package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          string `json:"ID" gorm:"primaryKey"`
	Email       string `json:"email" gorm:"unique"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Role        string `json:"role"`
	Department  string `json:"department"`
	Interests   string `json:"interests"`
	Description string `json:"description" gorm:"type:text"`
	Phone       string `json:"phone"`
	Password    []byte `json:"-"`
}
