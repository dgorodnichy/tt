package usecase

type DeleteProject struct {
	projects Remover
}

type Remover interface {
	Delete(id string) error
}

func NewDeleteProject(projects Remover) *DeleteProject {
	return &DeleteProject{projects}
}

func (dp *DeleteProject) Call(id string) (error) {
	err := dp.projects.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
