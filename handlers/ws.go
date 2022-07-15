package handlers

import (
	"github.com/gin-gonic/gin"
)

type Ws interface {
	Serve(c *gin.Context)
}
