package repositories

import "github.com/suttapak/cacf/models"

type ProductRepository interface {
	GetAll(shopID uint) ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Create(product models.Product) error
	Update(product models.Product) error
	Delete(id uint) error
}
