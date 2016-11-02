package main

import (
	"gopkg.in/validator.v2"
)

type User struct {
	Name     string `json:"name" validate:"nonzero"`
	Username string `json:"username" validate:"nonzero"`
}

func (s User) Validate() error {

	return validator.Validate(s)
}
