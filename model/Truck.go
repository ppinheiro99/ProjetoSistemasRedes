package model

import "github.com/jinzhu/gorm"

type Truck struct {
	gorm.Model `swaggerignore:"true"`
	Plate   string `gorm:"unique;not null" json:"plate"`
	Year   uint `gorm:"not null" json:"year"`
	Month uint `gorm:"not null" json:"month"`
	Km uint `gorm:"not null" json:"km"`
	Brand   string `gorm:"not null" json:"brand"`
	Latitude   float64 `gorm:"not null" json:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude"`
	
}