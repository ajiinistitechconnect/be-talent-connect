package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type MentorMenteeUsecase interface {
	BaseUsecase[model.MentorMentee]
}

type mentorMenteeUsecase struct {
	repo repository.MentorMenteeRepo
}

func (m *mentorMenteeUsecase) FindAll() ([]model.MentorMentee, error) {
	return m.repo.List()
}

func (m *mentorMenteeUsecase) FindById(id string) (*model.MentorMentee, error) {
	return m.repo.Get(id)
}

func (m *mentorMenteeUsecase) SaveData(payload *model.MentorMentee) error {
	return m.repo.Save(payload)
}

func (m *mentorMenteeUsecase) DeleteData(id string) error {
	return m.repo.Delete(id)
}

func NewMentorMenteeUsecase(repo repository.MentorMenteeRepo) MentorMenteeUsecase {
	return &mentorMenteeUsecase{
		repo: repo,
	}
}
