package router

import (
	"github.com/labstack/echo/v4"
	"github.com/leoncolt/conv/controller"
)

type Router struct {
	Group *echo.Group
	Conv  *controller.ConvController
}

func (router *Router) Register(echo *echo.Echo) {
	router.Group = echo.Group("/public")

	router.Conv = new(controller.ConvController)
	router.Conv.Register(router.Group)
}
