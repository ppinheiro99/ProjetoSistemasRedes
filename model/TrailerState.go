package model

import (
	"github.com/jinzhu/gorm"
)

type TrailerState struct {
	gorm.Model `swaggerignore:"true"`
	TrailerId uint `gorm:"not null" json:"trailer_id"`
	Latitude   float64 `gorm:"not null" json:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude"`
}