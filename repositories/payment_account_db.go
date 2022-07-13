package repositories

import (
	"github.com/suttapak/cacf/models"
	"gorm.io/gorm"
)

type paymentAccountRepository struct {
	db *gorm.DB
}

func NewPaymentAccountRepository(db *gorm.DB) PaymentAccoutRepository {
	return &paymentAccountRepository{db}
}

func (r paymentAccountRepository) GetAll(shopID uint) ([]models.PaymentAccount, error)
func (r paymentAccountRepository) GetByID(id uint) (*models.PaymentAccount, error)
func (r paymentAccountRepository) Create(paymentAccount *models.PaymentAccount) error
func (r paymentAccountRepository) Update(paymentAccount *models.PaymentAccount) error
func (r paymentAccountRepository) Delete(id uint) error
