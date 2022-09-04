package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

func (controller *UserController) Index(c *gin.Context) {
	users, err := controller.userService.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (controller *UserController) Create(c *gin.Context) {
	request := dto.UserRegistrationRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := controller.userService.Register(&request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *UserController) Get(c *gin.Context) {
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

	user, err := controller.userService.GetById(query.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) UpdatePassword(c *gin.Context) {
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

	user, err := controller.userService.UpdatePassword(query.ID, request.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserController) Update(c *gin.Context) {
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

	user, err := controller.userService.Update(query.ID, &request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
