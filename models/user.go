package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          string         `json:"ID" gorm:"primaryKey"`
	Email       string         `json:"email"`
	Name        string         `json:"name"`
	Picture     string         `json:"picture"`
	Role        string         `json:"role"`
	Stage       string         `json:"stage"`
	Department  string         `json:"department"`
	Interests   string         `json:"interests"`
	Description string         `json:"description" gorm:"type:text"`
	Phone       string         `json:"phone"`
	CreatedAt   gorm.DeletedAt `gorm:"column:created_at"`
	UpdatedAt   gorm.DeletedAt `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}
