package server

import (
	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/controllers"
)

func NewRouter(authMiddleware gin.HandlerFunc, adminAreaMiddleware gin.HandlerFunc, tournamentController *controllers.TournamentController, userController *controllers.UserController) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(authMiddleware)

	router.GET("/users", adminAreaMiddleware, userController.Index)
	router.POST("/users", adminAreaMiddleware, userController.Create)
	router.PUT("/users", adminAreaMiddleware, userController.Index)
	router.GET("/users/:id", adminAreaMiddleware, userController.Index)
	router.PUT("/users/:id/update_password", adminAreaMiddleware, userController.Index)

	router.GET("/tournaments", tournamentController.Index)
	router.GET("/tournaments/past", tournamentController.Index)
	router.GET("/tournaments/upcoming", tournamentController.Index)
	router.POST("/tournaments", tournamentController.Index)
	router.PUT("/tournaments/:id", tournamentController.Index)
	router.DELETE("/tournaments/:id", tournamentController.Index)

	return router
}
