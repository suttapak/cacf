package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type imageImageProductRepository struct {
	db *gorm.DB
}

func NewImageProductRepository(db *gorm.DB) ImageProductRepository {
	return &imageImageProductRepository{db}
}

func (r imageImageProductRepository) GetByProductID(productID uint) (*models.ImageProduct, error)
func (r imageImageProductRepository) CreateOrFirst(imageProduct models.ImageProduct) (*models.ImageProduct, error)
func (r imageImageProductRepository) Update(imageProduct models.ImageProduct) error
func (r imageImageProductRepository) Delete(id uint) error
