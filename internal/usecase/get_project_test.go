package usecase

import (
	"errors"
	"testing"

	memory "github.com/dgorodnichy/tt/internal/adapter/memory"
)

func TestGetProject(t *testing.T) {
	repo := memory.NewRepo()
	creater := NewCreateProject(repo)
	finder := NewGetProject(repo)

	_, err := creater.Call("1", "firstProjectName", "firstProjectDescription")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	got, err := finder.Call("1")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if got.Name() != "firstProjectName" {
		t.Errorf("incorrect name. got: %q, want: %q", got.Name(), "firstProjectName")
	}

	if got.Description() != "firstProjectDescription" {
		t.Errorf("incorrect description. got: %q, want: %q", got.Description(), "firstProjectDescription")
	}
}

func TestGetProjectNotFound(t *testing.T) {
	repo := memory.NewRepo()
	creater := NewCreateProject(repo)
	finder := NewGetProject(repo)

	_, err := creater.Call("1", "firstProjectName", "firstProjectDescription")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	got, err := finder.Call("2")
	if !errors.Is(err, memory.ErrProjectNotFound) {
		t.Fatalf("unexpected error: %q", err)
	}

	if got != nil {
		t.Error("expect nil")
	}
}
