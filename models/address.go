package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Address     string
	SubDistrict string
	District    string
	Province    string
	PostalCode  string
	CustomerID  uint
	Customer    Customer
	ShopID      uint
	Shop        Shop
}
