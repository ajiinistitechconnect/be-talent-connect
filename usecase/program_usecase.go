package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type ProgramUsecase interface {
	BaseUsecase[model.Program]
	BaseSearchUsecase[model.Program]
}

type programUsecase struct {
	repo repository.ProgramRepo
}

func (p *programUsecase) FindAll() ([]model.Program, error) {
	return p.repo.List()
}

func (p *programUsecase) FindById(id string) (*model.Program, error) {
	return p.repo.Get(id)
}

func (p *programUsecase) SaveData(payload *model.Program) error {
	return p.repo.Save(payload)
}

func (p *programUsecase) DeleteData(id string) error {
	return p.repo.Delete(id)
}

func (p *programUsecase) SearchBy(by map[string]any) ([]model.Program, error) {
	return p.repo.Search(by)
}

func NewProgramUsecase(repo repository.ProgramRepo) ProgramUsecase {
	return &programUsecase{
		repo: repo,
	}
}
