package services

import "github.com/suttapak/cacf/dto"

type ShopService interface {
	GetShop() (*dto.ShopReply, error)
	UpdateShop(dto.UpdateShopDTO) error
}
