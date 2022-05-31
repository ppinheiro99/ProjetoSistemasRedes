package model

import (
	"github.com/jinzhu/gorm"
)

type TravelMaps struct {
	gorm.Model `swaggerignore:"true"`
	TruckId int `gorm:"not null;"json:"truck_id"`
	TrailerId uint `json:"trailer_id"`
	Coords string `gorm:"not null" json:"coords"`
	Distance float64 `gorm:"not null" json:"distance"`
	Time string `gorm:"not null" json:"time"`
	StartCountry string `gorm:"not null" json:"start_country"`
	StartCity string `gorm:"not null" json:"start_city"`
	StartPostalCode string `gorm:"not null" json:"start_postal_code"`
	StartAddress string `gorm:"not null" json:"start_address"`
	EndCountry string `gorm:"not null" json:"end_country"`
	EndCity string `gorm:"not null" json:"end_city"`
	EndPostalCode string `gorm:"not null" json:"end_postal_code"`
	EndAddress string `gorm:"not null" json:"end_address"`
	TotalValue float64 `gorm:"not null" json:"total_value"`
	TotalLiters float64 `gorm:"not null" json:"total_liters"`
}