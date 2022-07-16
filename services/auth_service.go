package services

import (
	"errors"
	"strconv"

	fb "github.com/huandu/facebook/v2"
	"github.com/suttapak/cacf/dto"
	"github.com/suttapak/cacf/logs"
	"github.com/suttapak/cacf/models"
	"github.com/suttapak/cacf/repositories"
	"gorm.io/gorm"
)

type authService struct {
	shopRepo     repositories.ShopRespository
	facebookRepo repositories.FacebookRespository
	pageRepo     repositories.PageFacebookRepository
}

func NewAuthService(shopRepo repositories.ShopRespository,
	facebookRepo repositories.FacebookRespository, pageRepo repositories.PageFacebookRepository) AuthService {
	return &authService{shopRepo, facebookRepo, pageRepo}
}

func (s authService) SignIn(shop dto.SignInDTO) (*dto.SignInReply, error) {
	if _, err := s.shopRepo.Get(); err != nil {
		if err != gorm.ErrRecordNotFound {
			//TODO : handler error.
			return nil, err
		}
		//Save Shop to database.
		if err := s.shopRepo.Create(models.Shop{
			Model:      gorm.Model{ID: shop.ID},
			Name:       shop.Name,
			Email:      shop.Email,
			FacebookID: shop.ID,
			Facebook: models.Facebook{
				Model: gorm.Model{ID: shop.ID},
				Token: shop.UserToken,
				Name:  shop.Name,
				Email: shop.Email,
			},
		}); err != nil {
			//TODO : handler error.
			return nil, err
		}
	}
	//Update token.
	if err := s.facebookRepo.Update(models.Facebook{Model: gorm.Model{ID: shop.ID}, Token: shop.UserToken}); err != nil {
		//TODO : handler error.
		return nil, err
	}
	//Get page info form facebook.
	if result, err := fb.Get("/me/accounts", fb.Params{
		"access_token": shop.UserToken,
		"fields":       "access_token,name,id,picture{url}",
	}); err == nil {
		var res dto.SignInReply
		err = result.Decode(&res)
		if err != nil {
			//TODO : handler error.
			logs.Error(err)
			return nil, err
		}

		//Save page repo.
		//Check page repo.
		for _, pageInfo := range res.Data {
			id, err := strconv.Atoi(pageInfo.ID)
			if err != nil {
				//TODO : Handler error.
				return nil, err
			}
			if _, err := s.pageRepo.CreaetOrFind(models.PageFacebook{
				Model:       gorm.Model{ID: uint(id)},
				AccessToken: pageInfo.AccessToken,
				Name:        pageInfo.Name,
				Picture:     pageInfo.Picture.Data.Url,
				ShopID:      shop.ID,
			}); err != nil {
				//TODO : Handler error.
				return nil, err
			}

		}
		return &res, nil
	}
	return nil, errors.New("TODO : Handler error")
}
