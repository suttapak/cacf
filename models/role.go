package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name   string
	Code   string
	Admins []*Admin `gorm:"many2many:admin_roles;"`
}
