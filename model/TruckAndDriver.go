package model

import "github.com/jinzhu/gorm"

type TruckAndDriver struct {
	gorm.Model `swaggerignore:"true"`
	TruckId int `gorm:"not null;unique" json:"truck_id"`
	FirstDriverId   int `gorm:"not null" json:"first_driver_id"`
	SecondDriverId   int ` json:"second_driver_id"`
	
}