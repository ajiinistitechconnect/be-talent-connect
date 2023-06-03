package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type ParticipantUsecase interface {
	BaseUsecase[model.Participant]
}

type participantUsecase struct {
	repo    repository.ParticipantRepo
	user    UserUsecase
	program ProgramUsecase
}

func (m *participantUsecase) FindAll() ([]model.Participant, error) {
	return m.repo.List()
}

func (m *participantUsecase) FindById(id string) (*model.Participant, error) {
	return m.repo.Get(id)
}

func (m *participantUsecase) SaveData(payload *model.Participant) error {
	program, err := m.program.FindById(payload.ProgramID)
	if err != nil {
		return err
	}
	payload.Program = *program

	user, err := m.user.FindById(payload.UserID)
	if err != nil {
		return err
	}
	payload.User = *user

	return m.repo.Save(payload)
}

func (m *participantUsecase) DeleteData(id string) error {
	return m.repo.Delete(id)
}

func NewParticipantUsecase(
	repo repository.ParticipantRepo,
	user UserUsecase,
	program ProgramUsecase,
) ParticipantUsecase {
	return &participantUsecase{
		repo:    repo,
		user:    user,
		program: program,
	}
}
