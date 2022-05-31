package model

import (
	"github.com/jinzhu/gorm"
)

type DisplacementsAndSupply struct {
	gorm.Model `swaggerignore:"true"`
	SupplyId uint `gorm:"not null;"json:"supply_id"`
	DisplacementId uint `json:"displacement_id"`

}