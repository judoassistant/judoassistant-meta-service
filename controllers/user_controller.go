package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService}
}

func (tc *UserController) auth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (tc *UserController) index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
