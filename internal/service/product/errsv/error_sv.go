package errsv

// ErrService is an interface

type ErrProduct struct {
	Message string
}

func (e ErrProduct) Error() string {
	return e.Message
}

type ErrValidProduct struct {
	Message string
}

func (e ErrValidProduct) Error() string {
	return e.Message
}

type ErrNotFoundProduct struct {
	Message string
}

func (e ErrNotFoundProduct) Error() string {
	return e.Message
}

type ErrProductConflict struct {
	Message string
}

func (e ErrProductConflict) Error() string {
	return e.Message
}

type ErrInvalidRequest struct {
	Message string
}

func (e ErrInvalidRequest) Error() string {
	return e.Message
}
