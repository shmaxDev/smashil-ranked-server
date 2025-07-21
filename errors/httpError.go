package errors

import "fmt"

type HTTPError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *HTTPError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

func NewHttpError (status int, message string, err error) *HTTPError {
	return &HTTPError{
		StatusCode: status,
		Message: message,
		Err: err,
	}
}