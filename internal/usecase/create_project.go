package usecase

import (
	memory "github.com/dgorodnichy/tt/internal/adapter/memory"
	domain "github.com/dgorodnichy/tt/internal/domain"
)

type CreateProject struct {
	repo *memory.Repo
}

func NewCreateProject(repo *memory.Repo) *CreateProject {
	return &CreateProject{repo: repo}
}

func (uc *CreateProject) Call(id string, name string, description string) (*domain.Project, error) {
	project, err := domain.NewProject(id, name, description)
	if err != nil {
		return nil, err
	}
	uc.repo.Add(project)

	return project, nil
}
