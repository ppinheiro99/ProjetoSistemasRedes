package model

import "github.com/jinzhu/gorm"

type Trailer struct {
	gorm.Model `swaggerignore:"true"`
	Plate   string `gorm:"unique;not null" json:"plate"`
	Year   uint `gorm:"not null" json:"year"`
	Latitude   float64 `gorm:"not null" json:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude"`
}