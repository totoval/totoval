package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/totoval/framework/config"
	"github.com/totoval/framework/helpers"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/utils/crypt"
	"github.com/totoval/framework/utils/jwt"
	"totoval/app/http/requests"
	"totoval/app/models"
)

type Login struct {
	controller.BaseController
}

func (l *Login) Login(c *gin.Context) {
	// validate and assign requestData
	var requestData requests.UserLogin
	if !l.Validate(c, &requestData, true) {
		return
	}

	user := models.User{
		Email: &requestData.Email,
	}
	if err := m.H().First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": helpers.L(c, "auth.login.failed_not_exist")})
		return
	}

	if !crypt.BcryptCheck(*user.Password, requestData.Password) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": helpers.L(c, "auth.login.failed_wrong_password")})
		return
	}

	// create jwt
	newJwt := jwt.NewJWT(config.GetString("auth.sign_key"))
	username := ""
	if user.Name != nil {
		username = *user.Name
	}
	if token, err := newJwt.CreateToken(string(*user.ID), username); err == nil {
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": helpers.L(c, "auth.login.failed_token_generate_error")})
	return
}
