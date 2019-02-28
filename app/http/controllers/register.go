package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/model"
	"net/http"
	"totoval/app/http/requests"
	"totoval/app/models"
)

type Register struct {
	controller.BaseController
}

func (r *Register) Register(c *gin.Context) {
	// validate and assign requestData
	var requestData requests.UserRegister
	if !r.Validate(c, &requestData) {
		return
	}

	// determine if exist
	user := models.User{
		Email:    &requestData.Email,
	}
	if model.Exist(&user, true) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "user has been registered before"})
		return
	}

	// create user
	user.Password = &requestData.Password //@todo password encryption
	if err := model.Create(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "user register failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
	return
}
