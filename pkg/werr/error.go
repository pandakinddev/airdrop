package werr

import "fmt"

type wrappedError struct {
	msg string
	err error
}

// Unwrap ...
func (e wrappedError) Unwrap() error {
	return e.err
}

// Error ...
func (e wrappedError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.err)
}

// Wrap ...
func Wrap(err error, format string, args ...interface{}) error {
	return wrappedError{
		err: err,
		msg: fmt.Sprintf(format, args...),
	}
}
