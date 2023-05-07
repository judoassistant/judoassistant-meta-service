package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"go.uber.org/zap"
)

func NewAdminAreaMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(AuthUserKey).(*dto.UserResponseDTO)

		if !user.IsAdmin {
			logger.Info("Unable to authorize user for admin area")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
