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
	router.PUT("/users", adminAreaMiddleware, userController.Update)
	router.GET("/users/:id", adminAreaMiddleware, userController.Get)
	router.PUT("/users/:id/update_password", adminAreaMiddleware, userController.UpdatePassword)

	router.GET("/tournaments", tournamentController.Index)
	router.GET("/tournaments/past", tournamentController.GetPast)
	router.GET("/tournaments/upcoming", tournamentController.GetUpcoming)
	router.POST("/tournaments", tournamentController.Create)
	router.GET("/tournaments/:id", tournamentController.Get)
	router.PUT("/tournaments/:id", tournamentController.Update)
	router.DELETE("/tournaments/:id", tournamentController.Delete)

	return router
}
