package server

import (
	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/controllers"
)

func NewRouter() *gin.Engine {
  router := gin.New()
  router.Use(gin.Logger())

  tournament := new(controllers.TournamentController)

  router.GET("/tournaments", tournament.Get)

  return router
}

