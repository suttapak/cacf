package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Message    string
	Permission string
	ShopID     uint
	Shop       Shop
}
