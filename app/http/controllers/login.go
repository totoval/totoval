package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/helpers"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/model"
	"net/http"
	"totoval/app/http/requests"
	"totoval/app/models"
)

type Login struct {
	controller.BaseController
}

func (l *Login) Login(c *gin.Context) {
	// validate and assign requestData
	var requestData requests.UserLogin
	if !l.Validate(c, &requestData) {
		return
	}

	user := models.User{
		Email:    &requestData.Email,
		Password: &requestData.Password,
	}
	if err := model.First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error":  helpers.L(c, "user doesn't exist or password error")})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
	return
}
