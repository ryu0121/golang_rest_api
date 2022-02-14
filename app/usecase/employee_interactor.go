package usecase

import "golang_rest_api/app/domain"

type EmployeeInteractor struct {
	EmployeeRepository EmployeeRepository
}

// err がない場合もゼロ値が変える(名前付き戻り値の動作)
func (interactor *EmployeeInteractor) Add(e domain.Employee) (err error) {
	_, err = interactor.EmployeeRepository.Store(e)
	return
}

func (interactor *EmployeeInteractor) Employees() (employees domain.Employees, err error) {
	employees, err = interactor.EmployeeRepository.FindAll()
	return
}