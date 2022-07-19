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
	Carts       []*Cart  `gorm:"many2many:cart_products;"`
	Bills       []*Bill  `gorm:"many2many:bill_products;"`
	Orders      []*Order `gorm:"many2many:order_products;"`
}
