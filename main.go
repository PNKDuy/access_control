package main

import (
	"access_control/controller"
	_ "access_control/docs"
	"access_control/model"
	"access_control/model/message"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Access Control Service
// @description manipulate data from Kafka
// @version 1.0
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func main() {
	ctx := context.Background()
	go message.ConsumeMessageAndSyncDatabase(ctx)

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret")})

	model.ConnectToPostgresWithGorm()
	server := echo.New()

	server.POST("/access-control", controller.CheckPermission, jwtMiddleware)
	general := server.Group("/general")
	{
		general.POST("/:type", controller.Create, jwtMiddleware)
		general.GET("/:type", controller.Get, jwtMiddleware)
		general.GET("/:type/:id", controller.GetById, jwtMiddleware)
		general.PUT("/:type/:id", controller.Update, jwtMiddleware)
		general.DELETE("/:type/:id", controller.Delete, jwtMiddleware)
	}
	casbin := server.Group("/casbin")
	{
		casbin.GET("/:role", controller.GetCasbinByRole, jwtMiddleware)
	}
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	go server.Logger.Fatal(server.Start(":8081"))

}
