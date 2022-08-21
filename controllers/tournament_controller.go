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

func (tc *TournamentController) Index(c *gin.Context) {
	queryParams := dto.TournamentIndexQueryDTO{}
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournaments, err := tc.tournamentService.Get(queryParams.After, 10)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (tc *TournamentController) GetPast(c *gin.Context) {
	tournaments, err := tc.tournamentService.GetPast(10)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (tc *TournamentController) GetUpcoming(c *gin.Context) {
	tournaments, err := tc.tournamentService.GetPast(10)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (tc *TournamentController) Create(c *gin.Context) {
	request := dto.TournamentCreationRequestDTO{}
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := tc.tournamentService.Create(&request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (tc *TournamentController) Get(c *gin.Context) {

}

func (tc *TournamentController) Update(c *gin.Context) {}

func (tc *TournamentController) Delete(c *gin.Context) {}
