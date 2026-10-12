package domain

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var (
	ErrBlankID     = errors.New("blank id")
	ErrBlankName   = errors.New("blank name")
	ErrTooLongName = errors.New("too long name")
)

type Project struct {
	id          string
	name        string
	description string
}

func (p *Project) ID() string {
	return p.id
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

func validateID(id string) error {
	if id == "" {
		return ErrBlankID
	}
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

func NewProject(id string, name string, description string) (*Project, error) {
	name = normalizeName(name)
	err := validateName(name)
	if err != nil {
		return nil, err
	}
	err = validateID(id)
	if err != nil {
		return nil, err
	}

	return &Project{id: id, name: name, description: description}, nil
}
