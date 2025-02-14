package employee_storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/employee"
)

var ErrCardNumberExists = errors.New("card number already exists")

func NewEmployeeJSONFile(path string) *EmployeeJSONFile {
	return &EmployeeJSONFile{
		path:   path,
		lastId: -1,
	}
}

type EmployeeJSONFile struct {
	path   string
	lastId int
}

func (s *EmployeeJSONFile) Load() (employees map[int]models.Employee, err error) {
	file, err := os.Open(s.path)
	if err != nil {
		return
	}
	defer file.Close()

	var EmployeesJSON []dto.EmployeeDoc
	err = json.NewDecoder(file).Decode(&EmployeesJSON)
	if err != nil {
		return
	}

	employees = make(map[int]models.Employee)
	for _, em := range EmployeesJSON {
		employees[em.Id], err = dto.EmployeeDtoToModel(em)
	}

	return
}

func (s *EmployeeJSONFile) Save(employee models.Employee) (err error) {
	employees, err := s.Load()
	if err != nil {
		return
	}
	for _, value := range employees {
		if value.CardNumber == employee.CardNumber {
			err = ErrCardNumberExists
			return
		}
	}
	employeeList := make([]dto.EmployeeDoc, 0, len(employees))
	for _, e := range employees {
		employeeData := dto.EmployeeModelToDto(e)
		employeeList = append(employeeList, employeeData)
	}
	employeeList = append(employeeList, dto.EmployeeModelToDto(employee))

	// Convertir la lista en JSON
	data, err := json.MarshalIndent(employeeList, "", "  ")
	if err != nil {
		return err
	}

	// Escribir el JSON generado en el archivo
	err = os.WriteFile(s.path, data, 0644)
	s.lastId++
	return
}

func (s *EmployeeJSONFile) Erase(employee models.Employee) (err error) {
	employees, err := s.Load()
	if err != nil {
		return
	}
	employeeList := make([]dto.EmployeeDoc, 0, len(employees))
	for _, e := range employees {
		if e.Id.GetId() != employee.Id.GetId() {
			employeeList = append(employeeList, dto.EmployeeModelToDto(e))
		}
	}

	// Convertir la lista en JSON
	data, err := json.MarshalIndent(employeeList, "", "  ")
	if err != nil {
		return
	}

	// Escribir el JSON generado en el archivo
	err = os.WriteFile(s.path, data, 0644)
	return
}

func (s *EmployeeJSONFile) GetLastId() (id int, err error) {
	if s.lastId == -1 {
		employees, errLoad := s.Load()
		if errLoad != nil {
			return
		}
		id = employees[len(employees)].Id.GetId()
		s.lastId = id
		return
	}
	id = s.lastId
	return
}

func (s *EmployeeJSONFile) CheckCardNumber(cardNumber string) (err error) {
	employees, err := s.Load()
	if err != nil {
		return
	}
	for _, value := range employees {
		if value.CardNumber.GetCardNumber() == cardNumber {
			err = ErrCardNumberExists
			return
		}
	}
	return
}
