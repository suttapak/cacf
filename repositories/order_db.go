package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r orderRepository) GetAll(customerID uint) ([]models.Order, error) {
	panic("It was not imprement.")
}
func (r orderRepository) GetByID(id uint) (*models.Order, error) {
	panic("It was not imprement.")
}
func (r orderRepository) Create(order models.Order) error {
	panic("It was not imprement.")
}
func (r orderRepository) Update(order models.Order) error {
	panic("It was not imprement.")
}
func (r orderRepository) Delete(id uint) error {
	panic("It was not imprement.")
}
