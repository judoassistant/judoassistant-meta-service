package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/services"
)

type TournamentController struct {
	tournamentService *services.TournamentService
}

func NewTournamentController(tournamentService *services.TournamentService) *TournamentController {
	return &TournamentController{tournamentService}
}

type IndexQueryParamsDTO struct {
	After int `form:"after"`
}

func (tc *TournamentController) Index(c *gin.Context) {
	queryParams := IndexQueryParamsDTO{}
	if err := c.BindJSON(&queryParams); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// user := c.MustGet(middleware.AuthUserKey).(*dto.UserDTO)
	log.Println(queryParams.After)
}

func (tc *TournamentController) GetPast(c *gin.Context) {}

func (tc *TournamentController) GetUpcoming(c *gin.Context) {}

func (tc *TournamentController) Create(c *gin.Context) {}

func (tc *TournamentController) Get(c *gin.Context) {}

func (tc *TournamentController) Update(c *gin.Context) {}

func (tc *TournamentController) Delete(c *gin.Context) {}
