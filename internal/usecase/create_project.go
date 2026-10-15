package usecase

import (
	"github.com/dgorodnichy/tt/internal/domain"
)

type CreateProject struct {
	projects Adder
}

type Adder interface {
	Add(*domain.Project)
}

func NewCreateProject(projects Adder) *CreateProject {
	return &CreateProject{projects}
}

func (uc CreateProject) Call(id string, name string, description string) (*domain.Project, error) {
	project, err := domain.NewProject(id, name, description)
	if err != nil {
		return nil, err
	}
	uc.projects.Add(project)
	return project, nil
}
