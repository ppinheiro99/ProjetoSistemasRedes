package model

import (
	"github.com/jinzhu/gorm"
)

// Supply struct with fields that will be used in the database
// fields: id, liters, truckId, driverId
type Supply struct {
	gorm.Model `swaggerignore:"true"`
	Liters float64 `gorm:"not null" json:"liters"`
	Value float64 `gorm:"not null" json:"value"`
	TruckId string `gorm:"not null" json:"truck_id"`
	DriverId string `gorm:"not null" json:"driver_id"`
}

