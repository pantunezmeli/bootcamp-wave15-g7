package employee

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

func NewEmployeeSQL(database *sql.DB) *EmployeeSQL {
	return &EmployeeSQL{db: database}
}

type EmployeeSQL struct {
	db *sql.DB
}

func (r *EmployeeSQL) FindAll() (employees map[int]models.Employee, err error) {
	rows, err := r.db.Query("SELECT `id`, `id_card_number`, `first_name`, `last_name`, `warehouse_id` FROM employees")
	if err != nil {
		return
	}
	defer rows.Close()

	employees = make(map[int]models.Employee)
	var key int
	for rows.Next() {
		var employee models.Employee

		var id int
		var cardNumber string
		var firstName string
		var lastName string
		var warehouseId int

		err = rows.Scan(&id, &cardNumber, &firstName, &lastName, &warehouseId)
		if err != nil {
			return
		}

		employee.Id = value_objects.NewOptionalId(id)
		employee.CardNumber = value_objects.NewOptionalCardNumber(cardNumber)
		employee.FirstName = value_objects.NewOptionalName(firstName)
		employee.LastName = value_objects.NewOptionalName(lastName)
		employee.WarehouseId = value_objects.NewOptionalId(warehouseId)

		employees[key] = employee
		key++
	}

	err = rows.Err()
	return
}

func (r *EmployeeSQL) FindById(id int) (employee models.Employee, err error) {
	row := r.db.QueryRow("SELECT * FROM employees WHERE id = ?", id)
	if err = row.Err(); err != nil {
		return
	}

	var cardNumber string
	var firstName string
	var lastName string
	var warehouseId int

	err = row.Scan(&id, &cardNumber, &firstName, &lastName, &warehouseId)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrIdNotFound
	}

	employee.Id = value_objects.NewOptionalId(id)
	employee.CardNumber = value_objects.NewOptionalCardNumber(cardNumber)
	employee.FirstName = value_objects.NewOptionalName(firstName)
	employee.LastName = value_objects.NewOptionalName(lastName)
	employee.WarehouseId = value_objects.NewOptionalId(warehouseId)

	return
}

func (r *EmployeeSQL) New(employee models.Employee) (newEmployee models.Employee, err error) {
	result, err := r.db.Exec("INSERT INTO employees (id_card_number, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)", employee.CardNumber.GetCardNumber(), employee.FirstName.GetName(), employee.LastName.GetName(), employee.WarehouseId.GetId())
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok {
			if mysqlErr.Number == 1062 { // Duplicate entry (id_card_number is the only field that doesn't accept duplicates)
				err = ErrCardNumberNotUnique
				return
			}
			if mysqlErr.Number == 1452 { // A foregin key constraint fails (warehouse_id is the only ofreign key)
				err = ErrWarehouseIdNotFound
				return
			}
		}
		return
	}

	newEmployee = employee
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	newEmployee.Id = value_objects.NewOptionalId(int(id64))

	return
}

func (r *EmployeeSQL) Edit(id int, employee models.Employee) (updatedEmployee models.Employee, err error) {
	var employeeData []any

	query := "UPDATE employees SET "
	if employee.CardNumber.GetCardNumber() != "" {
		employeeData = append(employeeData, employee.CardNumber.GetCardNumber())
		query += "id_card_number = ?"
	}
	if employee.FirstName.GetName() != "" {
		employeeData = append(employeeData, employee.FirstName.GetName())
		query += ", first_name = ?"
	}
	if employee.LastName.GetName() != "" {
		employeeData = append(employeeData, employee.LastName.GetName())
		query += ", last_name = ?"
	}
	if employee.WarehouseId.GetId() != 0 {
		employeeData = append(employeeData, employee.WarehouseId.GetId())
		query += ", warehouse_id = ?"
	}
	query += " WHERE id = ?"
	employeeData = append(employeeData, id)

	result, err := r.db.Exec(query, employeeData...)
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok {
			if mysqlErr.Number == 1062 { // Duplicate entry (id_card_number is the only field that doesn't accept duplicates)
				err = ErrCardNumberNotUnique
				return
			}
			if mysqlErr.Number == 1452 { // A foregin key constraint fails (warehouse_id is the only foreign key)
				err = ErrWarehouseIdNotFound
				return
			}
		}
		return
	}

	if affectedRows, errRows := result.RowsAffected(); errRows != nil {
		err = errRows
		return
	} else {
		if affectedRows == 0 {
			fmt.Println("No rows affected")
			err = ErrIdNotFound
			return
		}
	}

	updatedEmployee, err = r.FindById(id)
	return
}

func (r *EmployeeSQL) DeleteById(id int) (err error) {
	result, err := r.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok {
			if mysqlErr.Number == 1451 { // A foregin key constraint fails (inbound_orders has foregin keys to employees)
				err = ErrInboundOrderFK
				return
			}
		}
		return
	}

	if affectedRows, errRows := result.RowsAffected(); errRows != nil {
		return
	} else {
		if affectedRows == 0 {
			err = ErrIdNotFound
			return
		}
	}

	return
}
