package dhcp

import (
	"errors"
)

var (
	ErrInvalidFormat  = errors.New("Invalid Format.")
	ErrDuplicateField = errors.New("Two fields with the same code found.")
)
