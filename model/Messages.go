package model

import "github.com/jinzhu/gorm"

type Messages struct {
	gorm.Model `swaggerignore:"true"`
	Message   string `gorm:"not null" json:"message"`
	SenderID   int `gorm:"not null" json:"sender_id"`
	ReceiverID   int `gorm:"not null" json:"receiver_id"`
}