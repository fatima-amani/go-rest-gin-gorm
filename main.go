package main

import (
	"go-gin-rest-gorm/config"
	"go-gin-rest-gorm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)

	router.Run(":8080")
}
