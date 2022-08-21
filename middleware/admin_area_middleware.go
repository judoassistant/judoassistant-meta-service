package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
)

func AdminAreaMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(AuthUserKey).(*dto.UserResponseDTO)

		if !user.IsAdmin {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
