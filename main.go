package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lalizita/go-crud-boilerplate/db"
	"github.com/lalizita/go-crud-boilerplate/internal/controller"
	"github.com/lalizita/go-crud-boilerplate/internal/repository"
	"github.com/lalizita/go-crud-boilerplate/internal/service"
	"github.com/lalizita/go-crud-boilerplate/routes"
)

func main() {
	e := echo.New()

	mongoClient, err := db.Connect()
	if err != nil {
		fmt.Println("Error connecting mongodb")
	}

	ctx := context.Background()
	e.GET("/healthCheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	consentCollection := mongoClient.Database("manager").Collection("consents")
	managerRepository := repository.NewTaskRepository(consentCollection)
	managerService := service.NewTaskService(ctx, managerRepository)
	managerController := controller.NewTaskController(managerService)
	managerRoutes := routes.NewManagerRoutes(managerController)
	managerRoutes.CreateRoutes(e)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
