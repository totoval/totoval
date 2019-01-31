package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"totoval/app/models"
	"totoval/routes"
)

func main() {
	cmd()
	//DB = &model.BaseModel{}
	//DB.Initialize()
	//
	//userObj := &model.User{}
	//DB.DB().Where("user_id = ?", 1).Find(userObj)
	//println(userObj.UserID)
	userObj := &models.User{}
	println(userObj.User().ID)

	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	router := &routes.Routes{Router: r}
	router.Register()

	r.Run()
}
