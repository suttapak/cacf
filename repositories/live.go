package repositories

import "github.com/suttapak/cacf/models"

type LiveRespository interface {
	GetAll() ([]models.Live, error)
	GetByID(id uint) (models.Live, error)
	Create(live models.Live) error
	Update(live models.Live) error
	Delete(id uint) error
}
