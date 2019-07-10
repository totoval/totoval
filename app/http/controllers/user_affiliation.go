package controllers

import (
	"net/http"

	"github.com/totoval/framework/helpers/toto"
	"github.com/totoval/framework/http/controller"
	"github.com/totoval/framework/request"
	"totoval/app/models"
)

type UserAffiliation struct {
	controller.BaseController
}

func (uaff *UserAffiliation) RenderAll(c *request.Context) {
	var u models.UserAffiliation
	c.HTML(http.StatusOK, "user_affiliation.nodes", toto.V{
		"data": u.All(),
	})

	return
}
