package memory

import (
	"errors"
	"slices"
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

	if !errors.Is(err, ErrProjectNotFound) {
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

func TestAllTwoProjects(t *testing.T) {
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

	projects := r.All()

	if len(projects) != 2 {
		t.Errorf("expected count: %v, got: %v", 2, len(projects))
	}

	projectNames := []string{}

	for _, p := range projects {
		projectNames = append(projectNames, p.Name())
	}

	expectedNames := []string{"firstProjectName", "secondProjectName"}

	for _, en := range expectedNames {
		if !slices.Contains(projectNames, en) {
			t.Errorf("Should contains: %v", en)
		}
	}
}

func TestAllNoProjects(t *testing.T) {
	r := NewRepo()
	projects := r.All()

	if len(projects) != 0 {
		t.Errorf("expected count: %v, got: %v", 0, len(projects))
	}
}

func TestDeleteProject(t *testing.T) {
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

	projects := r.All()

	if len(projects) != 2 {
		t.Fatalf("expected count: %v, got: %v", 2, len(projects))
	}

	err = r.Delete("1")
	if err != nil {
		t.Fatal(err)
	}

	projects = r.All()

	if len(projects) != 1 {
		t.Fatalf("expected count: %v, got: %v", 1, len(projects))
	}

	_, err = r.Get("1")
	if !errors.Is(err, ErrProjectNotFound) {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := r.Get("2")
	if err != nil {
		t.Fatal(err)
	}
	if got.Name() != "secondProjectName" {
		t.Fatal("unexpected project")
	}
}

func TestDeleteProjectNotFound(t *testing.T) {
	r := NewRepo()
	err := r.Delete("999")
	if !errors.Is(err, ErrProjectNotFound) {
		t.Fatalf("expect: %v, got: %v", ErrProjectNotFound, err)
	}
}
