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

//TODO : Imprement LiveRepository interface
func (r liveRepository) GetAll() ([]models.Live, error) {
	panic("It was not imprement.")
}
func (r liveRepository) GetByID(id uint) (*models.Live, error) {
	panic("It was not imprement.")
}
func (r liveRepository) Create(live models.Live) error {
	panic("It was not imprement.")
}
func (r liveRepository) Update(live models.Live) error {
	panic("It was not imprement.")
}
func (r liveRepository) Delete(id uint) error {
	panic("It was not imprement.")
}
