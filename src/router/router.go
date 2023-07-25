package router

import (
	"cake-store/src/model"

	"github.com/labstack/echo/v4"
)

type route struct {
	group          *echo.Group
	cakeController model.CakeController
}

func RouteService(group *echo.Group, cakeController model.CakeController) {
	rt := &route{
		group:          group,
		cakeController: cakeController,
	}
	rt.routerInit()
}

func (r *route) routerInit() {
	r.group.GET("/cakes", r.cakeController.HandleFindAll())
	r.group.POST("/cakes", r.cakeController.HandleCreate())
	r.group.GET("/cakes/:id", r.cakeController.HandleFindById())
	r.group.PUT("/cakes/:id", r.cakeController.HandleUpdate())
	r.group.DELETE("/cakes/:id", r.cakeController.HandleDelete())
}
