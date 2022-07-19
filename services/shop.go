package services

import "github.com/suttapak/cacf/dto"

type ShopService interface {
	GetShop(shopID uint) (*dto.ShopReply, error)
	UpdateShop(dto.UpdateShopDTO) error
}
