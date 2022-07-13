package models

import "gorm.io/gorm"

type PaymentAccount struct {
	gorm.Model
	Name          string
	AccountNumber string
	Bank          string
	ShopID        uint
	Shop          Shop
}
