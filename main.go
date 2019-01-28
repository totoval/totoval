package main

import (
	"github.com/gin-gonic/gin"
	routes "Wallet/routes"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	routes.Routes{Router: r}.Register()

	r.Run()
}
