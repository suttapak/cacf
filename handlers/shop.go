package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/cacf/dto"
	"github.com/suttapak/cacf/services"
	"github.com/suttapak/cacf/utilities"
)

type ShopHandler interface {
	GetShop(c *gin.Context)
	Update(c *gin.Context)
}

type shopHandler struct {
	shopService services.ShopService
}

func NewShopHandler(shopService services.ShopService) ShopHandler {
	return &shopHandler{shopService}
}

func (h shopHandler) GetShop(c *gin.Context) {
	shop, err := h.shopService.GetShop()
	if err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	c.JSON(utilities.HandlerReply(http.StatusOK, shop, c))
}
func (h shopHandler) Update(c *gin.Context) {
	shopIdString := c.Param("shopId")
	shopID, err := strconv.Atoi(shopIdString)
	if err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	shopDTO := dto.UpdateShopDTO{}
	if err := c.ShouldBindJSON(&shopDTO); err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}

	shopDTO.ID = uint(shopID)
	if err := h.shopService.UpdateShop(shopDTO); err != nil {
		c.AbortWithStatusJSON(utilities.HandlerReply(http.StatusBadRequest, err, c))
		return
	}
	c.JSON(utilities.HandlerReply(http.StatusOK, nil, c))
}
