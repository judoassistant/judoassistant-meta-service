package handler

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/errors"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"go.uber.org/zap"
)

var _shortNameRegex = regexp.MustCompile(`/^[a-z\_\-0-9]$/`)

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
	// Parse input
	queryParams := dto.TournamentIndexQueryDTO{}
	if err := c.ShouldBindUri(&queryParams); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validateShortName(queryParams.After); err != nil {
		return err
	}

	// List
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
	// Parse input
	request := &dto.TournamentCreationRequestDTO{}
	if err := c.ShouldBindJSON(request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validateShortName(request.ShortName); err != nil {
		return err
	}

	// Create
	user := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	response, err := handler.tournamentService.Create(request, user)
	if err != nil {
		return errors.Wrap(err, "unable to create tournament")
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (handler *tournamentHandler) Get(c *gin.Context) error {
	// Parse input
	query := &dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validateShortName(query.ShortName); err != nil {
		return err
	}

	// Get
	tournament, err := handler.tournamentService.GetByShortName(query.ShortName)
	if err != nil {
		return errors.Wrap(err, "unable to get tournament")
	}

	c.JSON(http.StatusOK, tournament)
	return nil
}

func (handler *tournamentHandler) Update(c *gin.Context) error {
	// Parse input
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(&query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	request := &dto.TournamentUpdateRequestDTO{}
	if err := c.ShouldBindJSON(request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validateShortName(request.ShortName); err != nil {
		return err
	}

	// Update
	user := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	tournament, err := handler.tournamentService.Update(query.ShortName, request, user)
	if err != nil {
		return errors.Wrap(err, "unable to update tournament")
	}

	c.JSON(http.StatusOK, tournament)
	return nil
}

func (handler *tournamentHandler) Delete(c *gin.Context) error {
	// Parse input
	query := dto.TournamentQueryDTO{}
	if err := c.ShouldBindUri(&query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validateShortName(query.ShortName); err != nil {
		return err
	}

	// Delete
	user := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if err := handler.tournamentService.Delete(query.ShortName, user); err != nil {
		return errors.Wrap(err, "unable to delete tournament")
	}

	return nil
}

func validateShortName(shortName string) error {
	if !_shortNameRegex.MatchString(shortName) {
		return errors.New("invalid shortName", errors.CodeBadRequest)
	}
	return nil
}
