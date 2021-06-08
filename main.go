package main

import (
	"access_control/controller"
	_ "access_control/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Access Control Service
// @description manipulate data from Kafka
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	server := echo.New()

	server.POST("/access-control", controller.CheckPermission)
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Logger.Fatal(server.Start(":8080"))
}
