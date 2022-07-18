package services

import "github.com/suttapak/cacf/dto"

type ProductService interface {
	GetAllProduct(shopID uint) ([]dto.ProductReply, error)
	GetOneProduct(productID uint) (*dto.ProductReply, error)
	CreateProduct(dto.CreateProductDTO) error
	UpdateProduct(dto.UpdateProductDTO) error
	DeleteProduct(productID uint) error
}
