package errordb

type ErrDB struct {
	Message string
}

func (e ErrDB) Error() string {
	return e.Message
}

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

type ErrConflict struct {
	Message string
}

func (e ErrConflict) Error() string {
	return e.Message
}

type ErrViolateFK struct {
	Message string
}

func (e ErrViolateFK) Error() string {
	return e.Message
}
