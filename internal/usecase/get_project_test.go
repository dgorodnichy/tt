package usecase

import (
	"errors"
	"testing"

	"github.com/dgorodnichy/tt/internal/adapter/memory"
	"github.com/dgorodnichy/tt/internal/domain"
)

var ErrNetwork = errors.New("network error")

type FakeRepo struct {
	project    *domain.Project
	err        error
	addCounter int
	getCounter int
	delCounter int
}

func NewFakeRepo(project *domain.Project, err error, addCounter int, getCounter int, delCounter int) *FakeRepo {
	return &FakeRepo{project, err, addCounter, getCounter, delCounter}
}

func (f *FakeRepo) Add(p *domain.Project) {
	f.addCounter++
}

func (f *FakeRepo) Get(id string) (*domain.Project, error) {
	f.getCounter++

	return f.project, f.err
}

func (f *FakeRepo) Delete(id string) error {
	f.delCounter++
	return f.err
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
	repo := NewFakeRepo(nil, ErrNetwork, 0, 0, 0)
	usecaseGetProject := NewGetProject(repo)
	got, err := usecaseGetProject.Call("1")
	if !errors.Is(err, ErrNetwork) {
		t.Fatalf("unexpected error: %q, want: %q", err, ErrNetwork)
	}

	if got != nil {
		t.Error("expect nil")
	}
}
