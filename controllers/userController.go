package controllers

import (
	"net/http"

	"github.com/Carmo-sousa/api/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}

var user models.User
var users models.Users

// Adiciona um novo usuário
func AddUser(c *gin.Context) {
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users.Add(&user)
	c.JSON(http.StatusCreated, &user)
}

func ShowUsers(c *gin.Context) {
	c.JSON(http.StatusCreated, &users)
}
