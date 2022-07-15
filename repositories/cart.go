package repositories

import "github.com/suttapak/cacf/models"

type CartRepository interface {
	GetAll(customerID uint) ([]models.Cart, error)
	GetByCustomerID(customerID uint) (*models.Cart, error)
	GetByID(id uint) (*models.Cart, error)
	Create(cart models.Cart) error
	Update(cart models.Cart) error
	Delete(id uint) error
}
