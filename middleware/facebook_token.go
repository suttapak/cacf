package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	fb "github.com/huandu/facebook/v2"
	"github.com/suttapak/cacf/logs"
)

func VerifyFacebookToken(c *gin.Context) {

	var token string
	if len(c.Request.Header["Authorization"]) <= 0 {
		if len(c.Query("Authorization")) >= 0 {
			token = c.Query("Authorization")
			fmt.Println("token", token)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, c.Request.Header["Authorization"])
			return
		}
	} else {
		token = strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	}

	result, err := fb.Get("/me", fb.Params{
		"access_token": token,
	})
	if err != nil {
		logs.Info("token: " + token)
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	i, err := strconv.Atoi(result.GetField("id").(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userID", uint(i))
}
