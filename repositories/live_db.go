package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type liveRepository struct {
	db *gorm.DB
}

func NewLiveRepository(db *gorm.DB) LiveRespository {
	return &liveRepository{db}
}

func (r liveRepository) GetAll() ([]models.Live, error)
func (r liveRepository) GetByID(id uint) (models.Live, error)
func (r liveRepository) Create(live models.Live) error
func (r liveRepository) Update(live models.Live) error
func (r liveRepository) Delete(id uint) error
