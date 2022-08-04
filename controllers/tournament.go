package controllers

import "github.com/gin-gonic/gin"
import "net/http"

type TournamentController struct{}

func (u TournamentController) Get(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

