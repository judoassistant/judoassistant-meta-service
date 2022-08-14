package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
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

func (controller *UserController) Update(c *gin.Context) {
	// TODO
}

func (controller *UserController) Get(c *gin.Context) {
	// TODO
}

func (controller *UserController) UpdatePassword(c *gin.Context) {
	// TODO
}
