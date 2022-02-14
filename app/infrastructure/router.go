package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang_rest_api/app/interfaces/controllers"
)

func Init() {
	router := echo.New()

	// Middleware
  router.Use(middleware.Logger())
  router.Use(middleware.Recover())

	employeeController := controllers.NewEmployeeController(NewSqlHandler())

	router.POST("/employees", func(c echo.Context) error { return employeeController.Create(c) })
	router.GET("/employees", func(c echo.Context) error { return employeeController.Index(c) })

	router.Logger.Fatal(router.Start(":8080"))
}