package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/totoval/framework/http/controller"
	"totoval/app/models"
)

type UserAffiliation struct {
	controller.BaseController
}

func (uaff *UserAffiliation) RenderAll(c *gin.Context) {
	var u models.UserAffiliation
	c.HTML(http.StatusOK, "user_affiliation.nodes", gin.H{
		"data": u.All(),
	})

	return
}
