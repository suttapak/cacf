package models

import "gorm.io/gorm"

type PageFacebook struct {
	gorm.Model
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	ShopID      uint
	Shop        Shop
}
