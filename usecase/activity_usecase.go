package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type ActivityUsecase interface {
	BaseUsecase[model.Activity]
}

type activityUsecase struct {
	repo    repository.ActivityRepo
	program ProgramUsecase
}

// DeleteData implements ActivityUsecase
func (a *activityUsecase) DeleteData(id string) error {
	return a.repo.Delete(id)
}

// FindAll implements ActivityUsecase
func (a *activityUsecase) FindAll() ([]model.Activity, error) {
	return a.repo.List()
}

// FindById implements ActivityUsecase
func (a *activityUsecase) FindById(id string) (*model.Activity, error) {
	return a.repo.Get(id)
}

// SaveData implements ActivityUsecase
func (a *activityUsecase) SaveData(payload *model.Activity) error {
	program, err := a.program.FindById(payload.ProgramID)
	if err != nil && program != nil {
		return err
	}
	return a.repo.Save(payload)
}

func NewActivityUsecase(repo repository.ActivityRepo, program ProgramUsecase) ActivityUsecase {
	return &activityUsecase{
		repo:    repo,
		program: program,
	}
}
