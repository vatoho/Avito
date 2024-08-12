package storage

import "errors"

var (
	ErrUserNotFound = errors.New("no users with such token")
)
