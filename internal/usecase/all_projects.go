package usecase

import (
	"github.com/dgorodnichy/tt/internal/domain"
)

type GetAll struct {
	projects Fetcher
}

type Fetcher interface {
	All() []*domain.Project
}

func NewAllProjects(projects Fetcher) *GetAll {
	return &GetAll{projects}
}

func (gall GetAll) Call() []*domain.Project {
	return gall.projects.All()
}
