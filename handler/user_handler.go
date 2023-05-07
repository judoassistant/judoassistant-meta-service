package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/service"
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
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService}
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
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := handler.userService.Register(&request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) Get(c *gin.Context) {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)

	if query.ID != authorizedUser.ID {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, err := handler.userService.GetById(query.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (handler *userHandler) UpdatePassword(c *gin.Context) {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request := dto.UserPasswordUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)

	if query.ID != authorizedUser.ID {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, err := handler.userService.UpdatePassword(query.ID, request.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (handler *userHandler) Update(c *gin.Context) {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request := dto.UserUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if query.ID != authorizedUser.ID {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, err := handler.userService.Update(query.ID, &request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
