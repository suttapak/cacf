package services

import "github.com/suttapak/cacf/dto"

type AddressService interface {
	GetAll(id uint) (dto.AddressReply, error)
}
