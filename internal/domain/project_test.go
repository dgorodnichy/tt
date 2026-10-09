package domain

import (
	"errors"
	"strings"
	"testing"
)

func TestNewProjectTrim(t *testing.T) {
	name := " \n\t firstProject "
	got, err := NewProject(name, "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got.name != "firstProject" {
		t.Errorf("Incorrect trim: '%s'. Expected: 'firstProject'", got.name)
	}
}

func TestNewProjectLength(t *testing.T) {
	name := strings.Repeat("a", 101)
	_, err := NewProject(name, "")

	if !errors.Is(err, ErrTooLongName) {
		t.Error("name length validation failed")
	}
}

func TestNewProjectBlankName(t *testing.T) {
	name := ""
	_, err := NewProject(name, "")
	if !errors.Is(err, ErrBlankName) {
		t.Error("name blank value validation failed")
	}
}

func TestNewProjectSpacesName(t *testing.T) {
	name := "    "
	_, err := NewProject(name, "")
	if !errors.Is(err, ErrBlankName) {
		t.Error("name blank value validation failed")
	}
}

func TestNewProjectCyrillicValid(t *testing.T) {
	name := strings.Repeat("ф", 100)
	_, err := NewProject(name, "")
	if err != nil {
		t.Error("cyrillic name length validation failed")
	}
}

func TestNewProjectCyrillicInvalid(t *testing.T) {
	name := strings.Repeat("ф", 101)
	_, err := NewProject(name, "")
	if !errors.Is(err, ErrTooLongName) {
		t.Error("cyrilic name length validation failed")
	}
}

func TestNewProject(t *testing.T) {
	name := "firstProject"
	description := "Description for the first project"

	got, err := NewProject(name, description)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if got.name != "firstProject" {
		t.Errorf("incorrect name: '%s'. Expected: 'firstProject'", got.name)
	}

	if got.description != "Description for the first project" {
		t.Errorf("incorrect description: '%s'. Expected: 'Description for the first project'", got.description)
	}
}
