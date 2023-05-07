package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/service"
)

type TournamentHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
	GetPast(c *gin.Context)
	GetUpcoming(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)
}

type tournamentHandler struct {
	tournamentService service.TournamentService
}

func NewTournamentHandler(tournamentService service.TournamentService) TournamentHandler {
	return &tournamentHandler{tournamentService}
}

func (handler *tournamentHandler) Index(c *gin.Context) {
	queryParams := dto.TournamentIndexQueryDTO{}
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournaments, err := handler.tournamentService.Get(queryParams.After, 10)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (handler *tournamentHandler) GetPast(c *gin.Context) {
	tournaments, err := handler.tournamentService.GetPast(10)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (handler *tournamentHandler) GetUpcoming(c *gin.Context) {
	tournaments, err := handler.tournamentService.GetUpcoming(10)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (handler *tournamentHandler) Create(c *gin.Context) {
	request := dto.TournamentCreationRequestDTO{}
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)

	response, err := handler.tournamentService.Create(user, &request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *tournamentHandler) Get(c *gin.Context) {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournament, err := handler.tournamentService.GetById(query.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournament)
}

func (handler *tournamentHandler) Update(c *gin.Context) {
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
	tournament, err := handler.tournamentService.Update(query.ID, &request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournament)
}

func (handler *tournamentHandler) Delete(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented) // TODO: Implement
}
