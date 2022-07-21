package services

import "github.com/suttapak/cacf/dto"

type AddressService interface {
	GetAllAddress(modelID uint) ([]dto.AddressReply, error)
	GetAddressByID(id uint) (*dto.AddressReply, error)
	CreateAddress(dto.CreateAddress) error
	UpdateAddreee(dto.UpdateAddress) error
	DeleteAddress(id uint) error
}
