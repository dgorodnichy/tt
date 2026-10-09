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

	if got.Name() != "firstProject" {
		t.Errorf("Incorrect trim: '%s'. Expected: 'firstProject'", got.Name())
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

func TestSetName(t *testing.T) {
	project, err := NewProject("oldName", "")
	if err != nil {
		t.Fatal("Unexpected error during create project")
	}

	err = project.SetName("newName")
	if err != nil {
		t.Fatal("Unexpected error during rename project")
	}

	if project.Name() != "newName" {
		t.Errorf("error during rename project: %v", project.Name())
	}
}

func TestSetNameBlank(t *testing.T) {
	project, err := NewProject("name", "")
	if err != nil {
		t.Fatal("Unexpected error during create project")
	}
	newName := " "
	err = project.SetName(newName)
	if !errors.Is(err, ErrBlankName) {
		t.Fatal("unexpected error during rename project")
	}

	if project.Name() != "name" {
		t.Error("name value should be initial")
	}
}

func TestSetNameTooLong(t *testing.T) {
	project, err := NewProject("name", "")
	if err != nil {
		t.Fatal("Unexpected error during create project")
	}
	newName := strings.Repeat("a", 101)
	err = project.SetName(newName)
	if !errors.Is(err, ErrTooLongName) {
		t.Fatal("too long validation failed")
	}

	if project.Name() != "name" {
		t.Error("name value should be initial")
	}
}

func TestSetNameTrim(t *testing.T) {
	project, err := NewProject("name", "")
	if err != nil {
		t.Fatal("Unexpected error during create project")
	}
	newName := " a "
	err = project.SetName(newName)
	if err != nil {
		t.Fatal("unexpected error during rename project")
	}
	if project.Name() != "a" {
		t.Error("incorrect trim during setting new name")
	}
}

func TestNewProject(t *testing.T) {
	name := "firstProject"
	description := "Description for the first project"

	got, err := NewProject(name, description)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Name() != "firstProject" {
		t.Fatalf("incorrect name: '%s'. Expected: 'firstProject'", got.Name())
	}

	if got.Description() != "Description for the first project" {
		t.Errorf("incorrect description: '%s'. Expected: 'Description for the first project'", got.Description())
	}
}
