package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"golang_rest_api/app/domain"
	"golang_rest_api/app/interfaces/database"
	"golang_rest_api/app/usecase"
)

type EmployeeController struct {
	Interactor usecase.EmployeeInteractor
}

func NewEmployeeController(sqlHandler database.SqlHandler) *EmployeeController {
	return &EmployeeController{
		Interactor: usecase.EmployeeInteractor{
			EmployeeRepository: &database.EmployeeRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *EmployeeController) Create(c echo.Context) (err error) {
	name := c.FormValue("Name")
	age, err := strconv.Atoi(c.FormValue("Age"))
	if err != nil {
		c.JSON(500, err)
		return
	}
	department := c.FormValue("Department")

	employee := domain.Employee{
		Name: name,
		Age: age,
		Department: department,
	}

	err = controller.Interactor.Add(employee)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(http.StatusCreated, employee)
	return
}

func (controller EmployeeController) Index(c echo.Context) (err error) {
	employees, err := controller.Interactor.Employees()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(http.StatusOK, employees)
	return
}