package server

import (
	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/controllers"
)

func NewRouter(authMiddleware gin.HandlerFunc, tournamentController *controllers.TournamentController, userController *controllers.UserController) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(authMiddleware)

	router.GET("/tournaments", tournamentController.Index)
	router.GET("/users/", userController.Index)
	router.POST("/users/register", userController.Create)

	return router
}
