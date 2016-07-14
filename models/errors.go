package models

import (
	"errors"
)

var (
	ErrRequiredMissing = errors.New("required attributes missing")
)
