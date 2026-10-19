package usecase

import (
	"slices"
	"testing"

	"github.com/dgorodnichy/tt/internal/adapter/memory"
)

func TestGetAllSuccess(t *testing.T) {
	repo := memory.NewRepo()
	creater := NewCreateProject(repo)
	fetcher := NewAllProjects(repo)

	_, err := creater.Call("1", "firstProjectName", "firstProjectDescription")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}
	_, err = creater.Call("2", "secondProjectName", "secondProjectDescription")
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}
	_, err = creater.Call("3", "thirdProjectName", "thirdProjectDescription")

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	projects := fetcher.Call()

	if len(projects) != 3 {
		t.Fatalf("Projects count. Got %v, Want: %v", len(projects), 3)
	}

	names := []string{}
	for _, ps := range projects {
		names = append(names, ps.Name())
	}

	expectedNames := []string{"firstProjectName", "secondProjectName", "thirdProjectName"}
	for _, en := range expectedNames {
		if !slices.Contains(names, en) {
			t.Errorf("should contains: %v", en)
		}
	}
}

func TestGetAllSuccessNoProjects(t *testing.T) {
	repo := memory.NewRepo()
	fetcher := NewAllProjects(repo)

	projects := fetcher.Call()

	if len(projects) != 0 {
		t.Errorf("Projects count. Got %v, Want: %v", len(projects), 0)
	}
}
