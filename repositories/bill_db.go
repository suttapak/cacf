package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type billRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) BillRepository {
	return &billRepository{db}
}

func (r billRepository) GetAll(customerID uint) ([]models.Bill, error) {
	panic("It was not imprement.")
}
func (r billRepository) GetByID(id uint) (*models.Bill, error) {
	panic("It was not imprement.")
}
func (r billRepository) Create(bill models.Bill) error {
	panic("It was not imprement.")
}
func (r billRepository) Update(bill models.Bill) error {
	panic("It was not imprement.")
}
func (r billRepository) Delete(id uint) error {
	panic("It was not imprement.")
}
