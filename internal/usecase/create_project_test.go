package usecase

import (
	"errors"
	"testing"

	memory "github.com/dgorodnichy/tt/internal/adapter/memory"
	"github.com/dgorodnichy/tt/internal/domain"
)

type FakeRepoAdd struct {
	AddCallCounter int
}

func (f *FakeRepoAdd) Add(p *domain.Project) {
	f.AddCallCounter++
}

func TestCreateProject(t *testing.T) {
	repo := memory.NewRepo()

	usecaseCreateProject := NewCreateProject(repo)
	got, err := usecaseCreateProject.Call("1", "firstProjectName", "firstProjectDesc")
	if err != nil {
		t.Fatal(err)
	}

	if got.Name() != "firstProjectName" {
		t.Errorf("incorrect name. Got: %q, want: %q", got.Name(), "firstProjectName")
	}

	if got.Description() != "firstProjectDesc" {
		t.Errorf("incorrect description. Got: %q, want: %q", got.Description(), "firstProjectDesc")
	}

	project, err := repo.Get(got.ID())
	if err != nil {
		t.Fatal(err)
	}

	if project.Name() != "firstProjectName" {
		t.Errorf("incorrect name. Got: %q, want: %q", project.Name(), "firstProjectName")
	}

	if project.Description() != "firstProjectDesc" {
		t.Errorf("incorrect description. Got: %q, want: %q", project.Description(), "firstProjectDesc")
	}
}

func TestCreateProjectBlankName(t *testing.T) {
	repo := memory.NewRepo()

	usecaseCreateProject := NewCreateProject(repo)
	_, err := usecaseCreateProject.Call("1", "", "firstProjectDesc")
	if !errors.Is(err, domain.ErrBlankName) {
		t.Fatalf("unexpected error: %q, want: %q", err, domain.ErrBlankName)
	}

	if _, err := repo.Get("1"); !errors.Is(err, memory.ErrProjectNotFound) {
		t.Errorf("got error: %q, want: %q", err, memory.ErrProjectNotFound)
	}
}

func TestCreateProjectBlankID(t *testing.T) {
	repo := memory.NewRepo()

	usecaseCreateProject := NewCreateProject(repo)
	got, err := usecaseCreateProject.Call("", "firstProjectName", "firstProjectDesc")
	if !errors.Is(err, domain.ErrBlankID) {
		t.Fatalf("unexpected error: %q, want: %q", err, domain.ErrBlankID)
	}

	if got != nil {
		t.Errorf("project should be nil. Got %s", got.Name())
	}
}

func TestCreateCallAdd(t *testing.T) {
	repo := &FakeRepoAdd{}
	usecaseCreateProject := NewCreateProject(repo)
	_, err := usecaseCreateProject.Call("1", "firstProjectName", "firstProjectDesc")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if repo.AddCallCounter != 1 {
		t.Errorf("got: %d, want %d", repo.AddCallCounter, 1)
	}
}

func TestCreateNotCallAdd(t *testing.T) {
	repo := &FakeRepoAdd{}
	usecaseCreateProject := NewCreateProject(repo)
	_, err := usecaseCreateProject.Call("", "firstProjectName", "firstProjectDesc")
	if !errors.Is(err, domain.ErrBlankID) {
		t.Errorf("unexpected error: %q, want: %q", err, domain.ErrBlankID)
	}

	if repo.AddCallCounter != 0 {
		t.Errorf("got: %d, want %d", repo.AddCallCounter, 0)
	}
}
