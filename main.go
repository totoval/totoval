package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/totoval/framework/config"
	"totoval/config"
	"totoval/routes"
)

func init() {
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
