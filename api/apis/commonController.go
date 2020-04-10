package apis

import (
	"errors"
)

var (
	errNotExist        = errors.New("not exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the  insertion")
	errUpdationFailed  = errors.New("Error in the  updation")
	errDeletionFailed  = errors.New("Error in the  deletion")
)
