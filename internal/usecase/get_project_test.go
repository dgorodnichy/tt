package usecase

import (
	"errors"
	"testing"

	"github.com/dgorodnichy/tt/internal/adapter/memory"
	"github.com/dgorodnichy/tt/internal/domain"
)

var ErrNetwork = errors.New("network error")

type FakeRepo struct {
}

func (f *FakeRepo) Get(id string) (*domain.Project, error) {
	return nil, ErrNetwork
}

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

func TestGetProjectRecieveNetworkError(t *testing.T) {
	repo := &FakeRepo{}
	usecaseGetProject := NewGetProject(repo)
	got, err := usecaseGetProject.Call("1")
	if !errors.Is(err, ErrNetwork) {
		t.Fatalf("unexpected error: %q, want: %q", err, ErrNetwork)
	}

	if got != nil {
		t.Error("expect nil")
	}
}
