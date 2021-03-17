package models

type SubjectType string

const (
	SubjectTypeUser SubjectType = "user"
)

type Subject interface {
	SubjectType() SubjectType
}
