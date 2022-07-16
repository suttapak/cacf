package repositories

import "github.com/suttapak/cacf/models"

type PageFacebookRepository interface {
	GetByID(uint) (*models.PageFacebook, error)
	Get() (*models.PageFacebook, error)
	GetAll() ([]models.PageFacebook, error)
	Create(models.PageFacebook) error
	Update(models.PageFacebook) error
	Delete(id uint) error
}
