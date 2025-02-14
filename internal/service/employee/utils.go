package employee

import dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/employee"

func validateFields(employeeData dto.EmployeeDoc) (err error) {
	if employeeData.CardNumber == "" || employeeData.FirstName == "" || employeeData.LastName == "" || employeeData.WarehouseId == 0 {
		err = ErrEmptyField
	}
	return

}
