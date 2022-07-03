package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name       string
	Email      string
	Summary    string
	FacebookID uint
	Facebook   Facebook
	Address    []Address
}
