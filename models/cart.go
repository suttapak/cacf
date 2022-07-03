package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CustomerID uint
	Customer   Customer
	Products   []Product
	Discount   float64
	Total      float64
}
