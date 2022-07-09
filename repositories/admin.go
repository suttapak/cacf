package repositories

import "github.com/suttapak/cacf/models"

type AdminRepository interface {
	Create(addmin models.Admin) error
	FindByEmail(email string) (*models.Admin, error)
	FindByID(id uint) (*models.Admin, error)
	FindAll() ([]models.Admin, error)
	Update(admin models.Admin) error
	Delete(id uint) error
}
