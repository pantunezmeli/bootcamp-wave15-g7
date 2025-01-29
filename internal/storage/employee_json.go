package storage

import (
	"encoding/json"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

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

func (s *EmployeeJSONFile) Load() (employees map[int]model.Employee, err error) {
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

	employees = make(map[int]model.Employee)
	for _, em := range EmployeesJSON {
		employees[em.Id], err = dto.EmployeeDtoToModel(em)
	}

	return
}

func (s *EmployeeJSONFile) Save(employee model.Employee) (err error) {
	employees, err := s.Load()
	if err != nil {
		return
	}
	employeeList := make([]dto.EmployeeDoc, 0, len(employees))
	for _, e := range employees {
		employeeList = append(employeeList, dto.EmployeeModelToDto(e))
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

func (s *EmployeeJSONFile) Erase(employee model.Employee) (err error) {
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
		return err
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
