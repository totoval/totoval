package main

import (
	"github.com/gin-gonic/gin"
	"totoval/config"
	"totoval/routes"
)

func main() {
	config.Initialize()

	//fmt.Println(c.Get("app.test"))
	//fmt.Println(c.Get("app.test1"))
	//test, _ := c.Get("app.test2").([]string)
	//fmt.Println(test[0])
	//fmt.Println(c.Get("app.test2"))
	//println("结束")


	//DB = &model.BaseModel{}
	//DB.Initialize()
	//
	//userObj := &models.User{}
	//println(userObj.User().ID)

	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	router := &routes.Routes{Router: r}
	router.Register()

	r.Run()
}
