package usecase

import "golang_rest_api/app/domain"


type EmployeeRepository interface {
    Store(domain.Employee) (int, error)
    FindAll() (domain.Employees, error)
}