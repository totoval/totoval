package main

import (
	"Wallet/app/models"
	"Wallet/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//DB = &models.BaseModel{}
	//DB.Initialize()
	//
	//userObj := &models.User{}
	//DB.DB().Where("user_id = ?", 1).Find(userObj)
	//println(userObj.UserID)
	userObj := &models.User{}
	println(userObj.User().UserID)

	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	router := &routes.Routes{Router: r}
	router.Register()

	r.Run()
}
