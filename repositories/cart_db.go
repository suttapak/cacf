package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r cartRepository) GetAll(customerID uint) ([]models.Cart, error) {
	panic("It was not imprement.")
}
func (r cartRepository) GetByID(id uint) (*models.Cart, error) {
	panic("It was not imprement.")
}
func (r cartRepository) GetByCustomerID(customerID uint) (*models.Cart, error) {
	panic("It was not imprement.")
}
func (r cartRepository) Create(cart models.Cart) error {
	panic("It was not imprement.")
}
func (r cartRepository) Update(cart models.Cart) error {
	panic("It was not imprement.")
}
func (r cartRepository) Delete(id uint) error {
	panic("It was not imprement.")
}
