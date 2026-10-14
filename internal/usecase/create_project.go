package usecase

import (
	domain "github.com/dgorodnichy/tt/internal/domain"
)

type CreateProject struct {
	repo ProjectsRepo
}

type ProjectsRepo interface {
	Add(p *domain.Project)
}

func NewCreateProject(repo ProjectsRepo) *CreateProject {
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
