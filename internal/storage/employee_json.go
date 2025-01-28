package loader

import (
	"encoding/json"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

func NewEmployeeJSONFile(path string) *EmployeeJSONFile {
	return &EmployeeJSONFile{
		path: path,
	}
}

type EmployeeJSONFile struct {
	path string
}

func (l *EmployeeJSONFile) Load() (e map[int]model.Employee, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var EmployeesJSON []dto.EmployeeDoc
	err = json.NewDecoder(file).Decode(&EmployeesJSON)
	if err != nil {
		return
	}

	e = make(map[int]model.Employee)
	for _, em := range EmployeesJSON {
		newId, errE := domain.NewId(em.Id)
		if errE != nil {
			err = errE
		}
		newCardNumber, errE := domain.NewCardNumber(em.CardNumber)
		if errE != nil {
			err = errE
		}
		newFirstName, errE := domain.NewName(em.FirstName)
		if errE != nil {
			err = errE
		}
		newLastName, errE := domain.NewName(em.LastName)
		if errE != nil {
			err = errE
		}
		newWarehouseId, errE := domain.NewId(em.WarehouseId)
		if errE != nil {
			err = errE
		}
		e[em.Id] = model.Employee{
			Id:          newId,
			CardNumber:  newCardNumber,
			FirstName:   newFirstName,
			LastName:    newLastName,
			WarehouseId: newWarehouseId,
		}
	}

	return
}

func (l *EmployeeJSONFile) Save(e map[int]model.Employee) error {
	EmployeeList := make([]dto.EmployeeDoc, 0, len(e))
	for _, employee := range e {
		EmployeeList = append(EmployeeList, dto.EmployeeDoc{
			Id:          employee.Id.GetId(),
			CardNumber:  employee.CardNumber.GetCardNumber(),
			FirstName:   employee.FirstName.GetName(),
			LastName:    employee.LastName.GetName(),
			WarehouseId: employee.WarehouseId.GetId(),
		})
	}

	// Convertir la lista en JSON
	data, err := json.MarshalIndent(EmployeeList, "", "  ")
	if err != nil {
		return err
	}

	// Escribir el JSON generado en el archivo
	err = os.WriteFile(l.path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
