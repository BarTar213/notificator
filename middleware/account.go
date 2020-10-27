package middleware

import (
	"github.com/BarTar213/notificator/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		account := models.Account{}
		err := c.ShouldBindHeader(&account)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, models.Response{Error: "invalid account headers"})
			return
		}
		c.Set("account", account)
		c.Next()
	}
}
