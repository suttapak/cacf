package dto

type CreateProductDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Count       int     `json:"count" binding:"required"`
	Code        string  `json:"code" binding:"required"`
	ShopID      uint    `json:"shop_id"`
}

type UpdateProductDTO struct {
	ID          uint    `json:"product_id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Count       int     `json:"count" binding:"required"`
	Code        string  `json:"code" binding:"required"`
	ShopID      uint    `json:"shop_id"`
}

type ProductReply struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Count       int     `json:"count" binding:"required"`
	Code        string  `json:"code" binding:"required"`
	ShopID      uint    `json:"shop_id"`
}
