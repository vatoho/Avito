package delivery

import "errors"

var ErrNoFieldsToUpdate = errors.New("update method must update at least 1 field")
