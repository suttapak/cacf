package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	CustomerID uint
	Customer   Customer
	Products   []*Product `gorm:"many2many:bill_products;"`
	Discount   float64
	Total      float64
	ExpressID  string
}
