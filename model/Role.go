package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model `swaggerignore:"true"`
	Name   string `gorm:"unique;not null" json:"name"`
}