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

func (p *Project) Name() string {
	return p.name
}

func (p *Project) Description() string {
	return p.description
}

func (p *Project) SetName(name string) error {
	normalizedName := normalizeName(name)
	if err := validateName(normalizedName); err != nil {
		return err
	}
	p.name = normalizedName
	return nil
}

func normalizeName(name string) string {
	return strings.TrimSpace(name)
}

func validateName(name string) error {
	if name == "" {
		return ErrBlankName
	}

	if utf8.RuneCountInString(name) > 100 {
		return ErrTooLongName
	}
	return nil
}

func NewProject(name string, description string) (*Project, error) {
	n := normalizeName(name)
	err := validateName(n)
	if err != nil {
		return nil, err
	}

	return &Project{name: n, description: description}, nil
}
