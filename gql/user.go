package gql

import (
	"github.com/wendorf/gqlgen-error/models"
)

//go:generate go run github.com/99designs/gqlgen

var _ models.Subject = (*User)(nil)

type User struct{}

func (u *User) SubjectType() models.SubjectType {
	return models.SubjectTypeUser
}
