package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
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
	query := dto.UserGetQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserDTO)

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
	request := dto.UserPasswordUpdateRequestDTO{}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserDTO)

	if request.ID != authorizedUser.ID {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, err := controller.userService.UpdatePassword(&request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserController) Update(c *gin.Context) {
	// TODO: Implement
	c.AbortWithStatus(http.StatusNotImplemented)
}
