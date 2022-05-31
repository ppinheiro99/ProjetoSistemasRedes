package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model `swaggerignore:"true"`
	Email   string `gorm:"unique;not null" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	RoleId uint `gorm:"not null" json:"role_id"`
	FirstName   string `gorm:"not null" json:"first_name"`
	LastName   string `gorm:"not null" json:"last_name"`
	Address   string `json:"address"`
	Country   string `json:"country"`
}