package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/config"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/errors"
	"github.com/judoassistant/judoassistant-meta-service/handler"
	"go.uber.org/zap"
)

func NewRouter(conf *config.Config, loggingMiddleware, authMiddleware, adminAreaMiddleware gin.HandlerFunc, tournamentHandler handler.TournamentHandler, userHandler handler.UserHandler, logger *zap.Logger) (*gin.Engine, error) {
	if err := setGinMode(conf.Environment); err != nil {
		return nil, err
	}

	router := gin.New()
	router.Use(loggingMiddleware)
	router.Use(gin.Recovery())

	// TODO: Add middleware
	router.GET("/users", wrapHandler(userHandler.Index, logger))
	router.POST("/users", wrapHandler(userHandler.Create, logger))
	router.POST("/users/authenticate", wrapHandler(userHandler.Authenticate, logger))
	router.PUT("/users/:id", wrapHandler(userHandler.Update, logger))
	router.GET("/users/:id", wrapHandler(userHandler.Get, logger))
	router.PUT("/users/:id/update_password", wrapHandler(userHandler.UpdatePassword, logger))

	router.GET("/tournaments", wrapHandler(tournamentHandler.Index, logger))
	router.GET("/tournaments/past", wrapHandler(tournamentHandler.ListPast, logger))
	router.GET("/tournaments/upcoming", wrapHandler(tournamentHandler.ListUpcoming, logger))
	router.POST("/tournaments", wrapHandler(tournamentHandler.Create, logger))
	router.GET("/tournaments/:id", wrapHandler(tournamentHandler.Get, logger))
	router.PUT("/tournaments/:id", wrapHandler(tournamentHandler.Update, logger))
	router.DELETE("/tournaments/:id", wrapHandler(tournamentHandler.Delete, logger))

	return router, nil
}

func setGinMode(environment config.Environment) error {
	switch environment {
	case config.EnvironmentProduction:
		gin.SetMode(gin.ReleaseMode)
	case config.EnvironmentDevelopment:
		gin.SetMode(gin.DebugMode)
	default:
		return errors.New(fmt.Sprintf("Unexpected environment %q", environment), errors.CodeInternal)
	}

	return nil
}

type handlerFunc func(c *gin.Context) error

func wrapHandler(handler handlerFunc, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err == nil {
			c.Status(errors.CodeOK)
			return
		}

		code := errors.Code(err)
		zapFields := []zap.Field{zap.Int("error_code", code), zap.Error(err)}
		if errors.IsServerSide(code) {
			logger.Error("Unable to handle request due to server-side error", zapFields...)
		} else {
			logger.Info("Unable to handle request due to client-side error", zapFields...)
		}

		body := &dto.ErrorResponseDTO{
			Code:    code,
			Message: err.Error(),
		}
		c.JSON(code, body)
	}
}
