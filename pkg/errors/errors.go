package errors

import "errors"

var (
	NOT_FOUND = errors.New("not found")
	ID_EXISTS = errors.New("id already exists")
)
