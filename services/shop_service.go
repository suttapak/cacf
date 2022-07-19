package services

import (
	"github.com/suttapak/cacf/dto"
	"github.com/suttapak/cacf/errs"
	"github.com/suttapak/cacf/logs"
	"github.com/suttapak/cacf/models"
	"github.com/suttapak/cacf/repositories"
	"gorm.io/gorm"
)

type shopService struct {
	shopRepo repositories.ShopRespository
}

func NewShopRepository(shopRepo repositories.ShopRespository) ShopService {
	return &shopService{shopRepo}
}

func (s shopService) GetShop(shopID uint) (*dto.ShopReply, error) {
	shop, err := s.shopRepo.Get()
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrorBadRequest
	}
	shopReply := dto.ShopReply{
		ID:         shop.ID,
		Name:       shop.Name,
		Email:      shop.Email,
		Summary:    shop.Summary,
		FacebookID: shop.FacebookID,
	}
	return &shopReply, nil
}
func (s shopService) UpdateShop(shopDTO dto.UpdateShopDTO) error {
	shopModel := models.Shop{
		Model:   gorm.Model{ID: shopDTO.ID},
		Name:    shopDTO.Name,
		Email:   shopDTO.Email,
		Summary: shopDTO.Summary,
	}
	if err := s.shopRepo.Update(shopModel); err != nil {
		logs.Error(err)
		return errs.ErrorBadRequest
	}
	return nil
}
