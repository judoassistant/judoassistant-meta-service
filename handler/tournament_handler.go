package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/service"
)

type TournamentHandler struct {
	tournamentService service.TournamentService
}

func NewTournamentHandler(tournamentService service.TournamentService) *TournamentHandler {
	return &TournamentHandler{tournamentService}
}

func (controller *TournamentHandler) Index(c *gin.Context) {
	queryParams := dto.TournamentIndexQueryDTO{}
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournaments, err := controller.tournamentService.Get(queryParams.After, 10)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (controller *TournamentHandler) GetPast(c *gin.Context) {
	tournaments, err := controller.tournamentService.GetPast(10)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (controller *TournamentHandler) GetUpcoming(c *gin.Context) {
	tournaments, err := controller.tournamentService.GetUpcoming(10)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (controller *TournamentHandler) Create(c *gin.Context) {
	request := dto.TournamentCreationRequestDTO{}
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)

	response, err := controller.tournamentService.Create(user, &request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *TournamentHandler) Get(c *gin.Context) {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournament, err := controller.tournamentService.GetById(query.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournament)
}

func (controller *TournamentHandler) Update(c *gin.Context) {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request := dto.TournamentUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Todo: Verify ownership
	tournament, err := controller.tournamentService.Update(query.ID, &request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournament)
}

func (controller *TournamentHandler) Delete(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented) // TODO: Implement
}
