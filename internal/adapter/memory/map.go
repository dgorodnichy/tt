package memory

import (
	"errors"

	domain "github.com/dgorodnichy/tt/internal/domain"
)

var ErrProjectNotFound = errors.New("project not found")

type Repo struct {
	projects map[string]*domain.Project
}

func NewRepo() *Repo {
	return &Repo{projects: make(map[string]*domain.Project)}
}

func (r *Repo) Add(p *domain.Project) {
	r.projects[p.ID()] = p
}

func (r *Repo) Get(id string) (*domain.Project, error) {
	if _, ok := r.projects[id]; !ok {
		return nil, ErrProjectNotFound
	}
	return r.projects[id], nil
}

func (r *Repo) All() []*domain.Project {
	projects := []*domain.Project{}
	for _, p := range r.projects {
		projects = append(projects, p)
	}

	return projects
}
