package repositories

import "github.com/suttapak/cacf/models"

type ImageProductRepository interface {
	GetByProductID(productID uint) (*models.ImageProduct, error)
	CreateOrFirst(imageProduct models.ImageProduct) (*models.ImageProduct, error)
	Update(imageProduct models.ImageProduct) error
	Delete(id uint) error
}
