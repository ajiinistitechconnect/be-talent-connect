package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type EvaluationUsecase interface {
	BaseUsecase[model.Evaluation]
}

type evaluationUsecase struct {
	repo        repository.EvaluationRepo
	user        UserUsecase
	participant ParticipantUsecase
}

// DeleteData implements EvaluationUsecase
func (e *evaluationUsecase) DeleteData(id string) error {
	return e.repo.Delete(id)
}

// FindAll implements EvaluationUsecase
func (e *evaluationUsecase) FindAll() ([]model.Evaluation, error) {
	return e.repo.List()
}

// FindById implements EvaluationUsecase
func (e *evaluationUsecase) FindById(id string) (*model.Evaluation, error) {
	return e.repo.Get(id)
}

// SaveData implements EvaluationUsecase
func (e *evaluationUsecase) SaveData(payload *model.Evaluation) error {
	panelist, err := e.user.FindById(payload.PanelistID)
	if err != nil {
		return err
	}
	payload.Panelist = *panelist
	_, err = e.participant.FindById(payload.ParticipantID)
	if err != nil {
		return err
	}
	// payload.Participant = *participant
	return e.repo.Save(payload)
}

func NewEvaluationUsecase(repo repository.EvaluationRepo) EvaluationUsecase {
	return &evaluationUsecase{
		repo: repo,
	}
}
