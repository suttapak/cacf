package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/cacf/dto"
	"github.com/suttapak/cacf/services"
	"github.com/suttapak/cacf/utilities"
)

type ProductHandler interface {
	GetAllProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandler {
	return &productHandler{productService: productService}
}

func (h productHandler) GetAllProduct(c *gin.Context) {
	shopIdString := c.Param("shopID")
	shopID, err := strconv.Atoi(shopIdString)
	if err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	product, err := h.productService.GetAllProduct(uint(shopID))
	if err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	c.JSON(utilities.HandlerReply(http.StatusOK, product, c))
}
func (h productHandler) GetProduct(c *gin.Context) {
	productIDString := c.Param("productID")
	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	product, err := h.productService.GetOneProduct(uint(productID))
	if err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	c.JSON(utilities.HandlerReply(http.StatusOK, product, c))
}
func (h productHandler) CreateProduct(c *gin.Context) {
	product := dto.CreateProductDTO{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	if err := h.productService.CreateProduct(product); err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
	}
	c.JSON(utilities.HandlerReply(http.StatusCreated, nil, c))
}
func (h productHandler) UpdateProduct(c *gin.Context) {
	product := dto.UpdateProductDTO{}
	productIdString := c.Param("productID")
	produtID, err := strconv.Atoi(productIdString)
	if err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	product.ID = uint(produtID)
	if err := h.productService.UpdateProduct(product); err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	c.JSON(utilities.HandlerReply(http.StatusOK, nil, c))
}
func (h productHandler) DeleteProduct(c *gin.Context) {
	productIdString := c.Param("productID")
	productID, err := strconv.Atoi(productIdString)
	if err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	if err := h.productService.DeleteProduct(uint(productID)); err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	c.JSON(utilities.HandlerReply(http.StatusOK, nil, c))
}
