package memory

import (
	"errors"
	"testing"

	domain "github.com/dgorodnichy/tt/internal/domain"
)

func TestAddSuccess(t *testing.T) {
	project, err := domain.NewProject("1", "firstProjectName", "firstProjectDesc")
	if err != nil {
		t.Fatalf("error during create project %v", err)
	}
	r := NewRepo()

	r.Add(project)

	got, err := r.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	if got.Name() != "firstProjectName" {
		t.Fatalf("incorrect name. Got %s, Want %s", got.Name(), "firstProjectName")
	}

	if got.Description() != "firstProjectDesc" {
		t.Errorf("incorrect description. Got %s, Want %s", got.Description(), "firstProjectDesc")
	}
}

func TestGetNotFound(t *testing.T) {
	project, err := domain.NewProject("1", "firstProjectName", "firstProjectDesc")
	if err != nil {
		t.Fatal(err)
	}
	r := NewRepo()
	r.Add(project)
	project, err = r.Get("2")

	if project != nil {
		t.Error("returns incorrect project")
	}

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("incorrect not found processing: %v", err)
	}
}

func TestGetFromTwoProjects(t *testing.T) {
	project1, err := domain.NewProject("1", "firstProjectName", "firstProjectDesc")
	if err != nil {
		t.Fatal(err)
	}
	project2, err := domain.NewProject("2", "secondProjectName", "secondProjectDesc")
	if err != nil {
		t.Fatal(err)
	}

	r := NewRepo()

	r.Add(project1)
	r.Add(project2)

	got, err := r.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	if got.Name() != "firstProjectName" {
		t.Fatalf("incorrect name. Got %s, Want %s", got.Name(), "firstProjectName")
	}

	if got.Description() != "firstProjectDesc" {
		t.Fatalf("incorrect description. Got %s, Want %s", got.Description(), "firstProjectDesc")
	}

	got, err = r.Get("2")
	if err != nil {
		t.Fatal(err)
	}

	if got.Name() != "secondProjectName" {
		t.Fatalf("incorrect name. Got %s, Want %s", got.Name(), "secondProjectName")
	}

	if got.Description() != "secondProjectDesc" {
		t.Fatalf("incorrect description. Got %s, Want %s", got.Description(), "secondProjectDesc")
	}
}

func TestGetFromTwoProjectsWithSameID(t *testing.T) {
	project1, err := domain.NewProject("1", "firstProjectName", "firstProjectDesc")
	if err != nil {
		t.Fatal(err)
	}
	project2, err := domain.NewProject("1", "secondProjectName", "secondProjectDesc")
	if err != nil {
		t.Fatal(err)
	}

	r := NewRepo()

	r.Add(project1)
	r.Add(project2)

	got, err := r.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	if got.Name() != "secondProjectName" {
		t.Fatalf("incorrect name. Got %s, Want %s", got.Name(), "secondProjectName")
	}

	if got.Description() != "secondProjectDesc" {
		t.Fatalf("incorrect description. Got %s, Want %s", got.Description(), "secondProjectDesc")
	}
}
