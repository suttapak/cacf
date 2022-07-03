package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint
	Customer   Customer
	Status     bool
	Products   []Product
	Discount   float64
	Total      float64
}
