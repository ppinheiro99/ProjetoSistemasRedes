package model

import (
	"github.com/jinzhu/gorm"
)

type TrailerHistory struct {
	gorm.Model `swaggerignore:"true"`
	Latitude   float64 `gorm:"not null" json:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude"`
	TrailerId uint `gorm:"not null" json:"trailer_id"`
}