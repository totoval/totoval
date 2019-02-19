package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/model"
	"net/http"
	"totoval/app/http/requests"
	"totoval/app/models"
)

type Register struct{
	controller.BaseController
}

func (r *Register) Register(c *gin.Context) {
	// validate and assign requestData
	var requestData requests.UserRegister
    if !r.Validate(c, &requestData) {return}

	user := models.User{
		Email: &requestData.Email,
		Password: &requestData.Password,
	}
	/**
	requests.UserRegister struct must contains models.User struct field
	*/
	if err := model.Create(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
	return
}
