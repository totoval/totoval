package main

import (
	"Wallet/app/models"
	"Wallet/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	cmd()
	//DB = &models.BaseModel{}
	//DB.Initialize()
	//
	//userObj := &models.User{}
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

