package repositories

import "github.com/suttapak/cacf/models"

type CustomerRepository interface {
	GetAll(shopID uint) (models.Customer, error)
	GetByID(id uint) (models.Customer, error)
	Create(customer models.Customer) error
	Update(customer models.Customer) error
	Delete(id uint) error
}
