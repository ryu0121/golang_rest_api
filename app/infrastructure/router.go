package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang_rest_api/app/interfaces/controllers"
)

func Init() {
	router := echo.New()

	// Middlewareの文法がRailsのサーバーの基幹部に使用されている「Rack」に似てる
	// Railsの知識が生きた
  router.Use(middleware.Logger())
  router.Use(middleware.Recover())

	employeeController := controllers.NewEmployeeController(NewSqlHandler())

	router.POST("/employees", func(c echo.Context) error { return employeeController.Create(c) })
	router.GET("/employees", func(c echo.Context) error { return employeeController.Index(c) })

	router.Logger.Fatal(router.Start(":8080"))
}