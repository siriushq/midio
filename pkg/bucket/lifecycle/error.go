package lifecycle

import (
	"fmt"
)

// Error is the generic type for any error happening during tag
// parsing.
type Error struct {
	err error
}

// Errorf - formats according to a format specifier and returns
// the string as a value that satisfies error of type tagging.Error
func Errorf(format string, a ...interface{}) error {
	return Error{err: fmt.Errorf(format, a...)}
}

// Unwrap the internal error.
func (e Error) Unwrap() error { return e.err }

// Error 'error' compatible method.
func (e Error) Error() string {
	if e.err == nil {
		return "lifecycle: cause <nil>"
	}
	return e.err.Error()
}
