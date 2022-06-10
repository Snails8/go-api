package domain

import (
	"errors"
)

type UserError error

type UserErrors struct {
	Items []UserError
}

var (
	ErrName = UserError(errors.New("invalid user name"))
)

func (err UserErrors) Error() string {
	message := "error: invalid user"

	for _, item := range err.Items {
		message += "," + item.Error()
	}
	return message
}
