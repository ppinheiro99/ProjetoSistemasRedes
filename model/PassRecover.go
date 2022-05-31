package model

import "github.com/jinzhu/gorm"

type PassRecover struct {
	gorm.Model `swaggerignore:"true"`
	Email   string `gorm:"unique;not null" json:"email"`
	Token   string `gorm:"not null;unique" json:"token"`
}