package database

import "golang_rest_api/app/domain"

type EmployeeRepository struct {
	SqlHandler SqlHandler
}

// CRUD のうちCreate, READ(ALL) を実装

// error は組み込みのインターフェース
func (repo *EmployeeRepository) Store(e domain.Employee) (id int, err error) {
	// infrastructureの内部実装に癒着している
	// infrastructureはクリーンアーキテクチャの4層のうち一番外の層
	// interfaces 層が依存していい場所じゃない
	// interfaces 層内にインターフェースを作る
	result, err := repo.SqlHandler.Exec(
		"INSERT INTO employees (name, age, department) VALUES (?, ?, ?)", e.Name, e.Age, e.Department,
	)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = int(id64)
	return
}

func (repo *EmployeeRepository) FindAll() (employees domain.Employees, err error) {
	rows, err := repo.SqlHandler.Query("SELECT id, name, age, department FROM employees")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		var department string

		if err := rows.Scan(&id, &name, &age, &department); err != nil {
			// 一つエラーが起きても処理を止めずに最後の行まで行う
			continue
		}

		employee := domain.Employee{
			ID: id,
			Name: name,
			Age: age,
			Department: department,
		}
		employees = append(employees, employee)
	}
	return
}