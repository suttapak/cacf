package repositories

import "github.com/suttapak/cacf/models"

type OrderRepository interface {
	GetAll(customerID uint) ([]models.Order, error)
	GetByID(id uint) (*models.Order, error)
	Create(order models.Order) error
	Update(order models.Order) error
	Delete(id uint) error
}
