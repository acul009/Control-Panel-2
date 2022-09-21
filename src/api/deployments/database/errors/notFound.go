package errors

import "fmt"

type errNotFound struct {
	name string
}

type ErrNotFound interface {
	NotFound()
}

func (err *errNotFound) Error() string {
	return fmt.Sprintf("Deployment not found: %s", err.name)
}

func (err *errNotFound) NotFound() {
}

func MakeNotFoundError(name string) *errNotFound {
	return &errNotFound{
		name: name,
	}
}
