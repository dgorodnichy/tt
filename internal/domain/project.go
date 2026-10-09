package domain

import (
	"errors"
	"strings"
	"unicode/utf8"
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
	n := strings.TrimSpace(name)
	if n == "" {
		return nil, ErrBlankName
	}

	if utf8.RuneCountInString(n) > 100 {
		return nil, ErrTooLongName
	}

	return &Project{name: n, description: description}, nil
}
