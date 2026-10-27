package usecase

import (
	"errors"
	"testing"

	"github.com/dgorodnichy/tt/internal/adapter/memory"
	"github.com/dgorodnichy/tt/internal/domain"
)

func TestRenameProjectSuccess(t *testing.T) {
	repo := memory.NewRepo()
	creater := NewCreateProject(repo)
	finder := NewGetProject(repo)
	renamer := NewRenameProject(repo)

	_, err := creater.Call("1", "firstProjectName", "firstProjectDescription")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = renamer.Call("1", "Renamed!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := finder.Call("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name() != "Renamed!" {
		t.Errorf("Want: %v, got: %v", "Renamed!", got.Name())
	}
}

func TestRenameProjectNotFound(t *testing.T) {
	repo := memory.NewRepo()
	renamer := NewRenameProject(repo)

	err := renamer.Call("1", "Renamed!")
	if !errors.Is(err, memory.ErrProjectNotFound) {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRenameProjectsFakeRepoBlankName(t *testing.T) {
	project, err := domain.NewProject("1", "name", "desc")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	repo := NewFakeRepo(project, nil, 0, 0, 0)
	renamer := NewRenameProject(repo)

	err = renamer.Call("1", "")
	if !errors.Is(err, domain.ErrBlankName) {
		t.Errorf("unexpected error: %v", err)
	}

	if repo.addCounter != 0 {
		t.Error("add was call")
	}
}

func TestRenameProjectFakeRepoSuccess(t *testing.T) {
	project, err := domain.NewProject("1", "name", "desc")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	repo := NewFakeRepo(project, nil, 0, 0, 0)
	renamer := NewRenameProject(repo)

	err = renamer.Call("1", "Renamed!")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if repo.addCounter != 1 {
		t.Error("add wasn't call")
	}
}
