package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	ShopID   uint
	Shop     Shop
	Roles    []*Role `gorm:"many2many:admin_roles;"`
}
