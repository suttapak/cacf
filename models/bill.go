package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	CustomerID uint
	Customer   Customer
	Products   []Product
	Discount   float64
	Total      float64
	ExpressID  string
}
