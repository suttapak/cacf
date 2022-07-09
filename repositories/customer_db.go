package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

//TODO : Imprement CustomerRepository interface
func (r customerRepository) GetAll(shopID uint) (models.Customer, error)
func (r customerRepository) GetByID(id uint) (models.Customer, error)
func (r customerRepository) Create(customer models.Customer) error
func (r customerRepository) Update(customer models.Customer) error
func (r customerRepository) Delete(id uint) error
