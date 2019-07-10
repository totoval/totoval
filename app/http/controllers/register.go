package controllers

import (
	"errors"
	"net/http"

	"github.com/totoval/framework/helpers/log"
	"github.com/totoval/framework/helpers/toto"
	"github.com/totoval/framework/request"

	"github.com/totoval/framework/hub"

	"github.com/totoval/framework/config"
	"github.com/totoval/framework/helpers"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/model/helper"
	"github.com/totoval/framework/utils/crypt"
	"github.com/totoval/framework/utils/jwt"

	"totoval/app/events"
	pbs "totoval/app/events/protocol_buffers"
	"totoval/app/http/requests"
	"totoval/app/models"
)

type Register struct {
	controller.BaseController
}

func (r *Register) Register(c *request.Context) {
	// validate and assign requestData
	var requestData requests.UserRegister
	if !r.Validate(c, &requestData, true) {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			responseErr, ok := err.(error)
			if ok {
				c.JSON(http.StatusUnprocessableEntity, toto.V{"error": responseErr.Error()})
				return
			}
			panic(err)
		}
	}()

	var token string
	var userId uint
	m.Transaction(func(TransactionHelper *helper.Helper) {
		// determine if exist
		user := models.User{
			Email: &requestData.Email,
		}
		if TransactionHelper.Exist(&user, true) {
			panic(errors.New(helpers.L(c, "auth.register.failed_existed")))
		}

		// create user
		// encrypt password //@todo move to model setter later
		encryptedPassword := crypt.Bcrypt(requestData.Password)
		user.Password = &encryptedPassword
		if err := TransactionHelper.Create(&user); err != nil {
			panic(errors.New(helpers.L(c, "auth.register.failed_system_error")))
		}

		// create jwt
		newJwt := jwt.NewJWT(config.GetString("auth.sign_key"))
		username := ""
		if user.Name != nil {
			username = *user.Name
		}
		var err error
		token, err = newJwt.CreateToken(string(*user.ID), username)
		if err != nil {
			panic(helpers.L(c, "auth.register.failed_token_generate_error"))
		}

		userId = *user.ID
	}, 1)

	// emit user-registered event
	ur := events.UserRegistered{}
	param := &pbs.UserRegistered{
		UserId:              uint32(userId),
		AffiliationFromCode: "",
	}
	if requestData.AffiliationFromCode != nil {
		param.AffiliationFromCode = *requestData.AffiliationFromCode
	}
	ur.SetParam(param)
	if errs := hub.Emit(&ur); errs != nil {
		log.Info("user registered event emit failed", toto.V{"event": ur, "errors": errs})
	}

	c.JSON(http.StatusOK, toto.V{"token": token})
	return
}
