package domain

import (
	"errors"
	"strings"
)

var (
	ErrBlankName   = errors.New("blank name")
	ErrTooLongName = errors.New("too long name")
)

type Project struct {
	name        string
	description string
}

func NewProject(name string, description string) (*Project, error) {
	n := strings.Trim(name, " ")
	if n == "" {
		return nil, ErrBlankName
	}

	if len([]rune(n)) >= 100 {
		return nil, ErrTooLongName
	}

	return &Project{name: n, description: description}, nil
}
