package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"go.uber.org/zap"
)

const AuthUserKey = "user"

func NewBasicAuthMiddleware(userService service.UserService, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, password, hasAuth := c.Request.BasicAuth()

		if !hasAuth {
			logger.Info("Received request without basic auth")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := userService.Authenticate(&dto.UserAuthenticationRequestDTO{
			Email:    email,
			Password: password,
		})

		if err != nil {
			logger.Info("Unable to authenticate user", zap.Error(err))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(AuthUserKey, user)
		c.Next()
	}
}
