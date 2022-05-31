package model

import "github.com/jinzhu/gorm"

type Locations struct {
	gorm.Model `swaggerignore:"true"`
	Name   string `gorm:"unique;not null" json:"name"`
	Latitude   float64 `gorm:"not null;unique" json:"latitude"`
	Longitude   float64 `gorm:"not null;unique" json:"longitude"`
}