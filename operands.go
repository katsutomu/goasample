package main

import (
	"github.com/goadesign/goa"
	"goa_sample/app"
	"strconv"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	sum := ctx.Left + ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sum)))
}

// Subtract runs the subtract action.
func (c *OperandsController) Subtract(ctx *app.SubtractOperandsContext) error {
	sub := ctx.Left - ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sub)))
}
