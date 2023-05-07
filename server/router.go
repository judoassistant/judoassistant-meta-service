package server

import (
	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/config"
	"github.com/judoassistant/judoassistant-meta-service/handler"
	"github.com/pkg/errors"
)

func newGinRouter(environment config.Environment) (*gin.Engine, error) {
	// Setup gin environment
	if environment == config.EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	} else if environment == config.EnvironmentDevelopment {
		gin.SetMode(gin.DebugMode)
	} else {
		return nil, errors.Errorf("Unexpected environment %q", environment)
	}

	// Setup router
	router := gin.New()
	router.Use(gin.Logger())
	return router, nil
}

func NewRouter(conf *config.Config, authMiddleware gin.HandlerFunc, adminAreaMiddleware gin.HandlerFunc, tournamentHandler handler.TournamentHandler, userHandler handler.UserHandler) (*gin.Engine, error) {
	router, err := newGinRouter(conf.Environment)
	if err != nil {
		return nil, err
	}

	router.Use(authMiddleware)

	router.GET("/users", adminAreaMiddleware, userHandler.Index)
	router.POST("/users", adminAreaMiddleware, userHandler.Create)
	router.PUT("/users", adminAreaMiddleware, userHandler.Update)
	router.GET("/users/:id", adminAreaMiddleware, userHandler.Get)
	router.PUT("/users/:id/update_password", adminAreaMiddleware, userHandler.UpdatePassword)

	router.GET("/tournaments", tournamentHandler.Index)
	router.GET("/tournaments/past", tournamentHandler.GetPast)
	router.GET("/tournaments/upcoming", tournamentHandler.GetUpcoming)
	router.POST("/tournaments", tournamentHandler.Create)
	router.GET("/tournaments/:id", tournamentHandler.Get)
	router.PUT("/tournaments/:id", tournamentHandler.Update)
	router.DELETE("/tournaments/:id", tournamentHandler.Delete)

	return router, nil
}
