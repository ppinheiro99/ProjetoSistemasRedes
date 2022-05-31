package model

import (
	"github.com/jinzhu/gorm"
)

type TruckHistory struct {
	gorm.Model `swaggerignore:"true"`
	Latitude   float64 `gorm:"not null" json:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude"`
	TruckId uint `gorm:"not null" json:"truck_id"`
	Rpm uint `gorm:"not null" json:"rpm"`
	Speed uint `gorm:"not null" json:"speed"`
}