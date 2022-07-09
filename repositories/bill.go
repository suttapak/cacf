package repositories

import "github.com/suttapak/cacf/models"

type BillRepository interface {
	GetAll(customerID uint) ([]models.Bill, error)
	GetByID(id uint) (*models.Bill, error)
	Create(bill models.Bill) error
	Update(bill models.Bill) error
	Delete(id uint) error
}
