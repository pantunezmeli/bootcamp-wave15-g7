package employee

type EmployeeRepository interface {
	// TODO
	FindAll() (err error)
	FindById() (err error)
	New() (err error)
	Update() (err error)
	DeleteById() (err error)
}
