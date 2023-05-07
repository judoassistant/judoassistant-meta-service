package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"go.uber.org/zap"
)

type UserHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)
	UpdatePassword(c *gin.Context)
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

func (handler *userHandler) Index(c *gin.Context) {
	users, err := handler.userService.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (handler *userHandler) Create(c *gin.Context) {
	request := dto.UserRegistrationRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		handler.logger.Info("Unable to map user registration request", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := handler.userService.Register(&request)
	if err != nil {
		handler.logger.Warn("Unable to register user", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) Get(c *gin.Context) {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		handler.logger.Info("Unable to map user get request", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)

	if query.ID != authorizedUser.ID {
		handler.logger.Info("Unable to authorize user get request")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, err := handler.userService.GetById(query.ID)
	if err != nil {
		handler.logger.Warn("Unable to get user", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (handler *userHandler) UpdatePassword(c *gin.Context) {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		handler.logger.Info("Unable to map update password request", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request := dto.UserPasswordUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		handler.logger.Info("Unable to map update password request", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)

	if query.ID != authorizedUser.ID {
		handler.logger.Info("Unable to authorize update password request")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, err := handler.userService.UpdatePassword(query.ID, request.Password)
	if err != nil {
		handler.logger.Warn("Unable to update password")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (handler *userHandler) Update(c *gin.Context) {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		handler.logger.Info("Unable to authorize user get request")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request := dto.UserUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		handler.logger.Info("Unable to authorize user get request")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if query.ID != authorizedUser.ID {
		handler.logger.Info("Unable to authorize user get request")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, err := handler.userService.Update(query.ID, &request)
	if err != nil {
		handler.logger.Warn("Unable to update user")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
