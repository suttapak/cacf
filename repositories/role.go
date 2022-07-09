package repositories

import "github.com/suttapak/cacf/models"

type RoleRepository interface {
	GetAll() ([]models.Role, error)
	GetByID(id uint) (*models.Role, error)
	Create(role models.Role) error
	Update(role models.Role) error
	Delete(id uint) error
}
