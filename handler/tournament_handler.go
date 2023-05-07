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
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
	ListPast(c *gin.Context)
	ListUpcoming(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)
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

func (handler *tournamentHandler) Index(c *gin.Context) {
	queryParams := dto.TournamentIndexQueryDTO{}
	if err := c.ShouldBindUri(&queryParams); err != nil {
		handler.logger.Info("Unable to map tournament index query", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournaments, err := handler.tournamentService.List(queryParams.After, 10)
	if err != nil {
		handler.logger.Warn("Unable to list tournaments", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (handler *tournamentHandler) ListPast(c *gin.Context) {
	tournaments, err := handler.tournamentService.ListPast(10)
	if err != nil {
		handler.logger.Warn("Unable to list past tournaments", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (handler *tournamentHandler) ListUpcoming(c *gin.Context) {
	tournaments, err := handler.tournamentService.ListUpcoming(10)
	if err != nil {
		handler.logger.Warn("Unable to list upcoming tournaments", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func (handler *tournamentHandler) Create(c *gin.Context) {
	request := dto.TournamentCreationRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		handler.logger.Info("Unable map create tournament request", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)

	response, err := handler.tournamentService.Create(user, &request)
	if err != nil {
		handler.logger.Warn("Unable to create tournament", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *tournamentHandler) Get(c *gin.Context) {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(&query); err != nil {
		handler.logger.Info("Unable map get tournament request", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tournament, err := handler.tournamentService.GetByID(query.ID)
	if err != nil {
		handler.logger.Warn("Unable to get tournament", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournament)
}

func (handler *tournamentHandler) Update(c *gin.Context) {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(&query); err != nil {
		handler.logger.Info("Unable to map query tournament request", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request := dto.TournamentUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		handler.logger.Info("Unable to map update tournament request", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// TODO: Authorize resource-level
	tournament, err := handler.tournamentService.Update(query.ID, &request)
	if err != nil {
		handler.logger.Warn("Unable to update tournament", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tournament)
}

func (handler *tournamentHandler) Delete(c *gin.Context) {
	err := handler.delete(c)
	if err == nil {
		c.Status(http.StatusOK)
		return
	}

	handler.logger.Warn("Error handling request", zap.Error(err))
	status := http.StatusInternalServerError
	if codedErr, ok := err.(errors.Coder); ok {
		status = codedErr.Code()
	}

	body := &dto.ErrorResponseDTO{
		Message: err.Error(),
	}
	c.JSON(status, body)

}

func (handler *tournamentHandler) delete(c *gin.Context) error {
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(&query); err != nil {
		return errors.WrapCode(err, "unable to map query tournament request", errors.CodeBadRequest)
	}

	// TODO: Authorize resource-level
	if err := handler.tournamentService.Delete(query.ID); err != nil {
		return errors.Wrap(err, "unable to delete tournament")
	}

	return nil
}
