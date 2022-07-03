package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Count       int
	Code        string
	ShopID      uint
	Carts       []Cart
}
