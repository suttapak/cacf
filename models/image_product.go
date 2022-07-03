package models

import "gorm.io/gorm"

type ImageProduct struct {
	gorm.Model
	ProductID uint
	Name      string
	Path      string
	Product   Product
}
