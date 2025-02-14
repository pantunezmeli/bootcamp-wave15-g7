package errsv

// ErrService is an interface

type ErrService struct {
	Message string
}

func (e ErrService) Error() string {
	return e.Message
}

type ErrValidEntity struct {
	Message string
}

func (e ErrValidEntity) Error() string {
	return e.Message
}

type ErrNotFoundEntity struct {
	Message string
}

func (e ErrNotFoundEntity) Error() string {
	return e.Message
}

type ErrConflict struct {
	Message string
}

func (e ErrConflict) Error() string {
	return e.Message
}

type ErrInvalidRequest struct {
	Message string
}

func (e ErrInvalidRequest) Error() string {
	return e.Message
}
