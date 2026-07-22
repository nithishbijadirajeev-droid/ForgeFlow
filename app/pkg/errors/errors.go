package errors

type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func New(status int, message string, err error) *AppError {
	return &AppError{
		StatusCode: status,
		Message:    message,
		Err:        err,
	}
}
