package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type pageFacebookRepository struct {
	db *gorm.DB
}

func NewPageFacebookRepository(db *gorm.DB) PageFacebookRepository {
	return &pageFacebookRepository{db}
}

func (r pageFacebookRepository) GetByID(uint) (*models.PageFacebook, error)
func (r pageFacebookRepository) Get() (*models.PageFacebook, error)
func (r pageFacebookRepository) GetAll() ([]models.PageFacebook, error)
func (r pageFacebookRepository) Create(models.PageFacebook) error
func (r pageFacebookRepository) CreaetOrFind(models.PageFacebook) (*models.PageFacebook, error)
func (r pageFacebookRepository) Update(models.PageFacebook) error
func (r pageFacebookRepository) Delete(id uint) error
