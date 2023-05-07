package server

import (
	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/config"
	"github.com/judoassistant/judoassistant-meta-service/handler"
	"github.com/pkg/errors"
)

func setGinMode(environment config.Environment) error {
	switch environment {
	case config.EnvironmentProduction:
		gin.SetMode(gin.ReleaseMode)
	case config.EnvironmentDevelopment:
		gin.SetMode(gin.DebugMode)
	default:
		return errors.Errorf("Unexpected environment %q", environment)
	}

	return nil
}

func NewRouter(conf *config.Config, loggingMiddleware, authMiddleware, adminAreaMiddleware gin.HandlerFunc, tournamentHandler handler.TournamentHandler, userHandler handler.UserHandler) (*gin.Engine, error) {
	if err := setGinMode(conf.Environment); err != nil {
		return nil, err
	}

	router := gin.New()
	router.Use(loggingMiddleware)
	router.Use(gin.Recovery())

	router.Use(authMiddleware)

	router.GET("/users", adminAreaMiddleware, userHandler.Index)
	router.POST("/users", adminAreaMiddleware, userHandler.Create)
	router.PUT("/users", adminAreaMiddleware, userHandler.Update)
	router.GET("/users/:id", adminAreaMiddleware, userHandler.Get)
	router.PUT("/users/:id/update_password", adminAreaMiddleware, userHandler.UpdatePassword)

	router.GET("/tournaments", tournamentHandler.Index)
	router.GET("/tournaments/past", tournamentHandler.ListPast)
	router.GET("/tournaments/upcoming", tournamentHandler.ListUpcoming)
	router.POST("/tournaments", tournamentHandler.Create)
	router.GET("/tournaments/:id", tournamentHandler.Get)
	router.PUT("/tournaments/:id", tournamentHandler.Update)
	router.DELETE("/tournaments/:id", tournamentHandler.Delete)

	return router, nil
}
