package main

import (
	"github.com/leoncolt/conv/app"
	_ "github.com/leoncolt/conv/common"
	_ "github.com/leoncolt/conv/docs"            // docs is generated by Swag CLI, you have to import it.
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Libre office convert API
// @version 1.0
// @description Api to convert file using libreoffice

// @host localhost:8080
// @BasePath /public

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name authorization
func main() {
	application := app.New()
	application.Engine.GET("/swagger/*", echoSwagger.WrapHandler)
	application.Register()
	application.Start()
}
