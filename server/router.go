package server

import (
	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/controllers"
)

func NewRouter(tournamentController *controllers.TournamentController) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/tournaments", tournamentController.Get)

	return router
}
