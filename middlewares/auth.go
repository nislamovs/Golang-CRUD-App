package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang-crud-app/utils"
	"net/http"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is missing",
		})
	}

	userId, err := utils.VerifyJWT(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "not authorised",
		})
		return
	}

	c.Set("userId", int64(userId))
	c.Next()
}
