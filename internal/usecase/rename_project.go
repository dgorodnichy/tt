package usecase

import (
	"github.com/dgorodnichy/tt/internal/domain"
)

type RenameProject struct {
	projects Renamer
}

type Renamer interface {
	Get(id string) (*domain.Project, error)
	Add(p *domain.Project)
}

func NewRenameProject(repo Renamer) *RenameProject {
	return &RenameProject{projects: repo}
}

func (rp RenameProject) Call(id string, newName string) error {
	project, err := rp.projects.Get(id)
	if err != nil {
		return err
	}

	err = project.SetName(newName)
	if err != nil {
		return err
	}

	rp.projects.Add(project)

	return nil
}
