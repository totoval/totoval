package main

import (
	"github.com/gin-gonic/gin"
	"totoval/config"
	"totoval/routes"
	c "totoval-framework/config"
)

func init(){
	config.Initialize()
}

func main() {

	//@cautions do not write DB ops at here

	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	router := &routes.Routes{Router: r}
	router.Register()

	r.Run(":" + c.GetString("app.port"))
}
