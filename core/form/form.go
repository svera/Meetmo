package form

type Error struct {
	Field   string
	Message string
}

func (e *Error) Error() string { return e.Message }
