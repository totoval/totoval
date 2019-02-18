package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/model"
	"net/http"
	"totoval/app/http/requests"
	"totoval/app/models"
)

type Register struct{}

func (*Register) Register(c *gin.Context) {
	var requestData requests.UserRegister
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	/**
	requests.UserRegister struct must contains models.User struct field
	 */
	if err := model.Create(requestData, &user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
}
