package repositories

import "github.com/suttapak/cacf/models"

type PaymentAccoutRepository interface {
	GetAll(shopID uint) ([]models.PaymentAccount, error)
	GetByID(id uint) (*models.PaymentAccount, error)
	Create(paymentAccount *models.PaymentAccount) error
	Update(paymentAccount *models.PaymentAccount) error
	Delete(id uint) error
}
