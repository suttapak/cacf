package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r productRepository) GetAll(shopID uint) ([]models.Product, error)
func (r productRepository) GetByID(id uint) (*models.Product, error)
func (r productRepository) Create(product models.Product) error
func (r productRepository) Update(product models.Product) error
func (r productRepository) Delete(id uint) error
