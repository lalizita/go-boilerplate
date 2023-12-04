package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lalizita/go-crud-boilerplate/internal/entity"
	"github.com/lalizita/go-crud-boilerplate/internal/service"
)

type TaskController struct {
	ManagerService service.ITaskService
}

func NewTaskController(m service.ITaskService) *TaskController {
	return &TaskController{ManagerService: m}
}

func (c *TaskController) CreateConsent(ec echo.Context) error {
	var input entity.TaskDTOInput
	if err := ec.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("Task is correct: %v\n", input)
	ctx := context.Background()
	ok, err := c.ManagerService.CreateConsent(ctx, input)
	if err != nil {
		return ec.String(http.StatusInternalServerError, "error create task in db")
	}
	return ec.JSON(http.StatusCreated, ok)
}
