package usecase

import (
	"errors"
	"testing"

	"github.com/dgorodnichy/tt/internal/adapter/memory"
)

func TestDeleteProject(t *testing.T) {
	repo := memory.NewRepo()
	creater := NewCreateProject(repo)
	finder := NewGetProject(repo)
	remover := NewDeleteProject(repo)
	fetcher := NewAllProjects(repo)

	_, err := creater.Call("1", "firstProjectName", "firstProjectDescription")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	_, err = creater.Call("2", "secondProjectName", "secondProjectDescription")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	projects := fetcher.Call()

	if len(projects) != 2 {
		t.Fatalf("Projects count. Got %v, Want: %v", len(projects), 2)
	}

	err = remover.Call("1")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	projects = fetcher.Call()

	if len(projects) != 1 {
		t.Fatalf("Projects count. Got %v, Want: %v", len(projects), 1)
	}

	_, err = finder.Call("1")
	if !errors.Is(err, memory.ErrProjectNotFound) {
		t.Fatalf("unexpected error: %q", err)
	}

	_, err = finder.Call("2")
	if err != nil {
		t.Fatalf("project %q not found", "2")
	}
}

func TestDeleteProjectRecieveNetworkError(t *testing.T) {
	repo := NewFakeRepo(nil, ErrNetwork, 0, 0, 0)
	usecaseDeleteProject := NewDeleteProject(repo)
	err := usecaseDeleteProject.Call("1")
	if !errors.Is(err, ErrNetwork) {
		t.Fatalf("unexpected error: %q, want: %q", err, ErrNetwork)
	}
}
