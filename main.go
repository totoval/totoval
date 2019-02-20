package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/totoval/framework/config"
	"github.com/totoval/framework/model"
	"totoval/config"
	"totoval/routes"
)

func init() {
	config.Initialize()
	model.Initialize()
}

func main() {

	//@cautions do not write DB ops at here

	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	routes.Register(r)

	r.Run(":" + c.GetString("app.port"))
}
