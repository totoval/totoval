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

type Register struct {
	controller.BaseController
}

func (r *Register) Register(c *gin.Context) {
	// validate and assign requestData
	var requestData requests.UserRegister
	if !r.Validate(c, &requestData, true) {
		return
	}

	// determine if exist
	user := models.User{
		Email: &requestData.Email,
	}
	if m.H().Exist(&user, true) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": helpers.L(c, "auth.register.failed_existed")})
		return
	}

	// create user
	// encrypt password //@todo move to model setter later
	encryptedPassword := crypt.Bcrypt(requestData.Password)
	user.Password = &encryptedPassword
	if err := m.H().Create(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": helpers.L(c, "auth.register.failed_system_error")})
		return
	}

	// add user affiliation
	if config.GetBool("user_affiliation.enable") {
		uaffPtr := &models.UserAffiliation{
			UserID: user.ID,
		}
		var err error
		if requestData.AffiliationFromCode != nil {
			err = uaffPtr.InsertNode(&user, *requestData.AffiliationFromCode)
		} else {
			err = uaffPtr.InsertNode(&user)
		}
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": helpers.L(c, "auth.register.failed_system_error")})
			return
		}
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

	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": helpers.L(c, "auth.register.failed_token_generate_error")})
	return
}
