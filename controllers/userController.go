package controllers

import (
	"net/http"

	"github.com/Carmo-sousa/api/models"
	"github.com/gin-gonic/gin"
)

var users models.Users

// Adiciona um novo usuário
func AddUser(c *gin.Context) {
	var user models.User
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

// Lista todos os usuários
func ShowUsers(c *gin.Context) {
	c.JSON(http.StatusCreated, &users)
}

// Recebe um id de usuário como parametro e retorna o mesmo
func ShowUser(c *gin.Context) {
	var user models.User
	username := c.Param("name")

	user, err := users.GetByName(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	email := c.Param("email")
	user, err := users.Delete(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
