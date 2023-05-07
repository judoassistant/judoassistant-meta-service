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

type UserHandler interface {
	Create(c *gin.Context) error
	Get(c *gin.Context) error
	Index(c *gin.Context) error
	Update(c *gin.Context) error
	UpdatePassword(c *gin.Context) error
}

type userHandler struct {
	userService service.UserService
	logger      *zap.Logger
}

func NewUserHandler(userService service.UserService, logger *zap.Logger) UserHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}

func (handler *userHandler) Index(c *gin.Context) error {
	users, err := handler.userService.GetAll()
	if err != nil {
		return errors.Wrap(err, "unable to get users")
	}

	c.JSON(http.StatusOK, users)
	return nil
}

func (handler *userHandler) Create(c *gin.Context) error {
	request := dto.UserRegistrationRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	response, err := handler.userService.Register(&request)
	if err != nil {
		return errors.Wrap(err, "unable to register user")
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (handler *userHandler) Get(c *gin.Context) error {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if query.ID != authorizedUser.ID {
		return errors.New("cannot access other users", errors.CodeForbidden)
	}

	user, err := handler.userService.GetById(query.ID)
	if err != nil {
		return errors.Wrap(err, "unable to get user")
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (handler *userHandler) UpdatePassword(c *gin.Context) error {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	request := dto.UserPasswordUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if query.ID != authorizedUser.ID {
		return errors.New("cannot access other users", errors.CodeForbidden)
	}

	user, err := handler.userService.UpdatePassword(query.ID, request.Password)
	if err != nil {
		return errors.Wrap(err, "unable to update password")
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (handler *userHandler) Update(c *gin.Context) error {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	request := dto.UserUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if query.ID != authorizedUser.ID {
		return errors.New("cannot access other users", errors.CodeForbidden)
	}

	user, err := handler.userService.Update(query.ID, &request)
	if err != nil {
		return errors.Wrap(err, "unable to update user")
	}

	c.JSON(http.StatusOK, user)
	return nil
}
