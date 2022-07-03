package models

import "gorm.io/gorm"

type Live struct {
	gorm.Model
	Title string
	After string
}
