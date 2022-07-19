package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CustomerID uint
	Customer   Customer
	Products   []*Product `gorm:"many2many:cart_products;"`
	Discount   float64
	Total      float64
}
