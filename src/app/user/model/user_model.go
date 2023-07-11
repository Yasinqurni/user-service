package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"name"`
	Address string `gorm:"address"`
	Email   string `gorm:"email"`
	Phone   string `gorm:"phone"`
}
