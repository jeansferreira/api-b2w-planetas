package helpers

// errorString is a trivial implementation of error, check https://blog.golang.org/error-handling-and-go
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
func NewError(text string) error {
	return &errorString{text}
}
