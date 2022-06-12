package schema

import "fmt"

type WrappedError struct {
	Code    int
	Message string
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%v: %s", w.Code, w.Message)
}

func ErrorWrap(code int, message string) *WrappedError {
	return &WrappedError{
		Code:    code,
		Message: message,
	}
}
