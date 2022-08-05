package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/services"
)

type TournamentController struct {
	tournamentService *services.TournamentService
}

func NewTournamentController(tournamentService *services.TournamentService) *TournamentController {
	return &TournamentController{tournamentService}
}

func (tc *TournamentController) Get(c *gin.Context) {
	user := c.MustGet(middleware.AuthUserKey).(dto.UserDTO)
	log.Println(user.Email, user.ID, user.IsAdmin)
	response := dto.TournamentResponseBody{
		Name:     "Hello",
		Location: "Foo",
	}
	c.JSON(http.StatusOK, response)
}
