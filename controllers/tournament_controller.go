package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/services"
)

type TournamentController struct {
	tournamentService *services.TournamentService
}

func NewTournamentController(tournamentService *services.TournamentService) *TournamentController {
	return &TournamentController{tournamentService}
}

func (tc *TournamentController) Get(c *gin.Context) {
	response := dto.TournamentResponseBody{
		Name:     "Hello",
		Location: "Foo",
	}
	c.JSON(http.StatusOK, response)
}
