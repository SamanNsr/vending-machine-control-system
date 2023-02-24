package vendingMachineHttpController

import (
	"github.com/labstack/echo/v4"
	vendingMachineDomain "github.com/samannsr/vending-machine-control-system/internal/vending_machine/domain"
)

type Router struct {
	controller vendingMachineDomain.HttpController
}

func NewRouter(controller vendingMachineDomain.HttpController) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Register(e *echo.Group) {
	e.GET("/vending-machine/:id", r.controller.GetVendingMachineById)
	e.POST("/insert-coin", r.controller.InsertCoin)
	e.POST("/select-product", r.controller.SelectProduct)
}
