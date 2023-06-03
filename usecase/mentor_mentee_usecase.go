package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type MentorMenteeUsecase interface {
	BaseUsecase[model.MentorMentee]
}

type mentorMenteeUsecase struct {
	repo    repository.MentorMenteeRepo
	user    UserUsecase
	program ProgramUsecase
}

func (m *mentorMenteeUsecase) FindAll() ([]model.MentorMentee, error) {
	return m.repo.List()
}

func (m *mentorMenteeUsecase) FindById(id string) (*model.MentorMentee, error) {
	return m.repo.Get(id)
}

func (m *mentorMenteeUsecase) SaveData(payload *model.MentorMentee) error {
	program, err := m.program.FindById(payload.ProgramID)
	if err != nil {
		return err
	}
	payload.Program = *program

	mentor, err := m.user.FindById(payload.MentorID)
	if err != nil {
		return err
	}
	payload.Mentor = *mentor

	mentee, err := m.user.FindById(payload.ParticipantID)
	if err != nil {
		return err
	}
	payload.Participant = *mentee

	return m.repo.Save(payload)
}

func (m *mentorMenteeUsecase) DeleteData(id string) error {
	return m.repo.Delete(id)
}

func NewMentorMenteeUsecase(
	repo repository.MentorMenteeRepo,
	user UserUsecase,
	program ProgramUsecase,
) MentorMenteeUsecase {
	return &mentorMenteeUsecase{
		repo:    repo,
		user:    user,
		program: program,
	}
}
