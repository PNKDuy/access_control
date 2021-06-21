package main

import (
	"access_control/controller"
	_ "access_control/docs"
	"access_control/model"
	"access_control/model/request"
	"context"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Access Control Service
// @description manipulate data from Kafka
// @version 1.0
// @host localhost:8081
// @BasePath /
func main() {
	go func(){
	ctx := context.Background()
	request.ConsumeMessageAndProduceBack(ctx)
	}()

	model.ConnectToPostgresWithGorm()
	server := echo.New()

	server.POST("/access-control", controller.CheckPermission)
	general := server.Group("/general")
	{
		general.POST("/:type", controller.Create)
		general.GET("/:type", controller.Get)
		general.GET("/:type/:id", controller.GetById)
		general.PUT("/:type/:id", controller.Update)
		general.DELETE("/:type/:id", controller.Delete)
	}
	casbin := server.Group("/casbin")
	{
		casbin.GET("/:role", controller.GetCasbinByRole)
	}
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Logger.Fatal(server.Start(":8081"))
}
