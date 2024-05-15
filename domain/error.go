package domain

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredential = errors.New("invalid credential")
