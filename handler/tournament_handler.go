package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/errors"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"go.uber.org/zap"
)

type TournamentHandler interface {
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
	Get(c *gin.Context) error
	ListPast(c *gin.Context) error
	ListUpcoming(c *gin.Context) error
	Index(c *gin.Context) error
	Update(c *gin.Context) error
}

type tournamentHandler struct {
	tournamentService service.TournamentService
	logger            *zap.Logger
}

func NewTournamentHandler(tournamentService service.TournamentService, logger *zap.Logger) TournamentHandler {
	return &tournamentHandler{
		tournamentService: tournamentService,
		logger:            logger,
	}
}

func (handler *tournamentHandler) Index(c *gin.Context) error {
	queryParams := dto.TournamentIndexQueryDTO{}
	if err := c.ShouldBindUri(&queryParams); err != nil {
		return errors.WrapCode(err, "unable to map request", errors.CodeBadRequest)
	}

	tournaments, err := handler.tournamentService.List(queryParams.After, 10)
	if err != nil {
		return errors.Wrap(err, "unable to list tournaments")
	}

	c.JSON(http.StatusOK, tournaments)
	return nil
}

func (handler *tournamentHandler) ListPast(c *gin.Context) error {
	tournaments, err := handler.tournamentService.ListPast(10)
	if err != nil {
		return errors.Wrap(err, "unable to list past tournaments")
	}

	c.JSON(http.StatusOK, tournaments)
	return nil
}

func (handler *tournamentHandler) ListUpcoming(c *gin.Context) error {
	tournaments, err := handler.tournamentService.ListUpcoming(10)
	if err != nil {
		return errors.Wrap(err, "unable to list upcoming tournaments")
	}

	c.JSON(http.StatusOK, tournaments)
	return nil
}

func (handler *tournamentHandler) Create(c *gin.Context) error {
	request := dto.TournamentCreationRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapCode(err, "unable to map request", errors.CodeBadRequest)
	}

	user := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)

	response, err := handler.tournamentService.Create(user, &request)
	if err != nil {
		return errors.Wrap(err, "unable to create tournament")
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (handler *tournamentHandler) Get(c *gin.Context) error {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(&query); err != nil {
		return errors.WrapCode(err, "unable to map request", errors.CodeBadRequest)
	}

	tournament, err := handler.tournamentService.GetByID(query.ID)
	if err != nil {
		return errors.Wrap(err, "unable to get tournament")
	}

	c.JSON(http.StatusOK, tournament)
	return nil
}

func (handler *tournamentHandler) Update(c *gin.Context) error {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(&query); err != nil {
		return errors.WrapCode(err, "unable to map request", errors.CodeBadRequest)
	}

	request := dto.TournamentUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// TODO: Authorize resource-level
	tournament, err := handler.tournamentService.Update(query.ID, &request)
	if err != nil {
		return errors.Wrap(err, "unable to update tournament")
	}

	c.JSON(http.StatusOK, tournament)
	return nil
}

func (handler *tournamentHandler) Delete(c *gin.Context) error {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(&query); err != nil {
		return errors.WrapCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// TODO: Authorize resource-level
	if err := handler.tournamentService.Delete(query.ID); err != nil {
		return errors.Wrap(err, "unable to delete tournament")
	}

	return nil
}
