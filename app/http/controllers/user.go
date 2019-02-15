package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"totoval/app/models"
)

type User struct{}

func (*User) Info(c *gin.Context) {
	mUser := &models.User{}
	c.JSON(http.StatusOK, mUser.User())
}
