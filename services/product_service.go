package services

import (
	"github.com/suttapak/cacf/dto"
	"github.com/suttapak/cacf/errs"
	"github.com/suttapak/cacf/logs"
	"github.com/suttapak/cacf/models"
	"github.com/suttapak/cacf/repositories"
	"gorm.io/gorm"
)

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductRepository(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo}
}

func (s productService) GetAllProduct(shopID uint) ([]dto.ProductReply, error) {
	//Get products from database.
	products, err := s.GetAllProduct(shopID)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrorBadRequest
	}
	//Handler reply.
	productReply := []dto.ProductReply{}
	for _, product := range products {
		productReply = append(productReply, dto.ProductReply{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Count:       product.Count,
			Code:        product.Code,
			ShopID:      product.ShopID,
		})
	}
	return productReply, nil
}
func (s productService) GetOneProduct(productID uint) (*dto.ProductReply, error) {
	//Get Product of reciver database.
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrorBadRequest
	}
	//Handler reply.
	productReply := dto.ProductReply{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Count:       product.Count,
		Code:        product.Code,
		ShopID:      product.ShopID,
	}
	return &productReply, nil
}
func (s productService) CreateProduct(productDTO dto.CreateProductDTO) error {
	//Create product in database.
	if err := s.productRepo.Create(models.Product{
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		Count:       productDTO.Count,
		Code:        productDTO.Code,
		ShopID:      productDTO.ShopID,
	}); err != nil {
		logs.Error(err)
		return errs.ErrorBadRequest
	}
	//GetProduct to handler reply.
	return nil
}
func (s productService) UpdateProduct(productDTO dto.UpdateProductDTO) error {
	//Update product in database path s.productRepo.Update(productDTO.ID, productDTO)
	if err := s.productRepo.Update(models.Product{
		Model:       gorm.Model{ID: productDTO.ID},
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		Count:       productDTO.Count,
		Code:        productDTO.Code,
		ShopID:      productDTO.ShopID,
	}); err != nil {
		logs.Error(err)
		return errs.ErrorBadRequest
	}
	return nil
}
func (s productService) DeleteProduct(productID uint) error {
	//Delete product in database.
	if err := s.productRepo.Delete(productID); err != nil {
		logs.Error(err)
		return errs.ErrorBadRequest
	}
	return nil
}
