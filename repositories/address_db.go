package repositories

import (
	"errors"

	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r addressRepository) GetAll(models string, modelID, ID uint) (address []models.Address, err error) {
	switch models {
	case "shop":
		if err = r.db.Where("shop_id = ? and id = ?", modelID, ID).First(&address).Error; err != nil {
			return nil, err
		}
		return address, nil
	case "customer":
		if err = r.db.Where("customer_id = ? and id = ?", modelID, ID).First(&address).Error; err != nil {
			return nil, err
		}
		return address, nil
	}
	return nil, errors.New("error : model not found")
}
func (r addressRepository) GetByID(id uint) (models.Address, error)
func (r addressRepository) Create(address models.Address) error
func (r addressRepository) Update(address models.Address) error
func (r addressRepository) Delete(id uint) error
