package model

import (
	"github.com/jinzhu/gorm"
)

type TruckState struct {
	gorm.Model `swaggerignore:"true"`
	TruckId uint `gorm:"not null" json:"truck_id"`
	Latitude   float64 `gorm:"not null" json:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude"`
	Rpm uint `gorm:"not null" json:"rpm"`
	Speed uint `gorm:"not null" json:"speed"`
}