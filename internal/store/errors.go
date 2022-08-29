package store

import "errors"

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("record was not found")
	// ErrWrongPassword ...
	ErrWrongPassword = errors.New("wrong password")
)
