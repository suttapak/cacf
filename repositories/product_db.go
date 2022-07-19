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

func (r productRepository) GetAll() ([]models.Product, error) {
	panic("It was not imprement.")
}
func (r productRepository) GetByID(id uint) (*models.Product, error) {
	panic("It was not imprement.")
}
func (r productRepository) Create(product models.Product) error {
	panic("It was not imprement.")
}
func (r productRepository) Update(product models.Product) error {
	panic("It was not imprement.")
}
func (r productRepository) Delete(id uint) error {
	panic("It was not imprement.")
}
