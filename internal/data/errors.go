package data

import (
	"errors"
	
)

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrMismatchedHashAndPassword = errors.New("password is wrong")
var ErrDbConn = errors.New("database con failed")
var ErrInternalServerError = errors.New("internal server error")
