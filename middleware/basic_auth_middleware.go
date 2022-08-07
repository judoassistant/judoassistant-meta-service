package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/services"
)

const AuthUserKey = "user"

func BasicAuthMiddleware(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, password, hasAuth := c.Request.BasicAuth()

		if !hasAuth {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := userService.Authenticate(&dto.UserAuthenticationRequestDTO{
			Email:    email,
			Password: password,
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(AuthUserKey, user)
		c.Next()
	}
}
