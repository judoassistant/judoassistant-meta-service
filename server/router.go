package server

import (
	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/handler"
)

func NewRouter(authMiddleware gin.HandlerFunc, adminAreaMiddleware gin.HandlerFunc, tournamentHandler handler.TournamentHandler, userHandler handler.UserHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
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

	return router
}
