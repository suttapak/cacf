package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r messageRepository) Get(permission string) (*models.Message, error) {
	message := models.Message{}
	if err := r.db.Model(&message).Where("permission = ?", permission).First(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}
func (r messageRepository) Create(message models.Message) error {
	if err := r.db.Model(&message).FirstOrCreate(&message).Error; err != nil {
		return err
	}
	return nil
}
func (r messageRepository) Update(message models.Message) error {
	if err := r.db.Updates(&message).Error; err != nil {
		return err
	}
	return nil
}
func (r messageRepository) Delete(message models.Message) error {
	if err := r.db.Where("id = ?", message.ID).Delete(&message).Error; err != nil {
		return err
	}
	return nil
}
