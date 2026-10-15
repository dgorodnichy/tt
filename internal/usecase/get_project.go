package usecase

import (
	"github.com/dgorodnichy/tt/internal/domain"
)

type GetProject struct {
	projects Getter
}

type Getter interface {
	Get(id string) (*domain.Project, error)
}

func NewGetProject(projects Getter) *GetProject {
	return &GetProject{projects: projects}
}

func (uc GetProject) Call(id string) (*domain.Project, error) {
	project, err := uc.projects.Get(id)
	if err != nil {
		return nil, err
	}

	return project, nil
}
