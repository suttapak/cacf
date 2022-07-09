package repositories

import "github.com/suttapak/cacf/models"

type AddressRepository interface {
	GetAll(models string, modelID, ID uint) ([]models.Address, error)
	GetByID(id uint) (models.Address, error)
	Create(address models.Address) error
	Update(address models.Address) error
	Delete(id uint) error
}
