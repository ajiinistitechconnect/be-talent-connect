package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type MentoringScheduleUsecase interface {
	BaseUsecase[model.MentoringSchedule]
}

type mentoringScheduleUsecase struct {
	repo repository.MentoringScheduleRepo
}

func (m *mentoringScheduleUsecase) FindAll() ([]model.MentoringSchedule, error) {
	return m.repo.List()
}

func (m *mentoringScheduleUsecase) FindById(id string) (*model.MentoringSchedule, error) {
	return m.repo.Get(id)
}

func (m *mentoringScheduleUsecase) SaveData(payload *model.MentoringSchedule) error {
	return m.repo.Save(payload)
}

func (m *mentoringScheduleUsecase) DeleteData(id string) error {
	return m.repo.Delete(id)
}

func NewMentoringScheduleUsecase(repo repository.MentoringScheduleRepo) MentoringScheduleUsecase {
	return &mentoringScheduleUsecase{
		repo: repo,
	}
}
