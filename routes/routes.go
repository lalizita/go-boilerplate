package routes

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lalizita/go-crud-boilerplate/internal/controller"
)

type TaskRouteController struct {
	managerController *controller.TaskController
}

func NewManagerRoutes(m *controller.TaskController) TaskRouteController {
	return TaskRouteController{m}
}

func (r *TaskRouteController) CreateRoutes(e *echo.Echo) {
	log.Println("Creating routes for tasks...")
	e.POST("/create", r.managerController.CreateConsent)
}
