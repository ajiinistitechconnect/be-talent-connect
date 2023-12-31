package usecase

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type EvaluationUsecase interface {
	BaseUsecase[model.Evaluation]
	GetByProgramUser(program_id string, user_id string, stage string) ([]model.Evaluation, error)
	GetEvaluateeByProgramPanelist(program_id string, panelist_id string) ([]model.Evaluation, error)
}

type evaluationUsecase struct {
	repo           repository.EvaluationRepo
	user           UserUsecase
	participant    ParticipantUsecase
	questionAnswer QuestionAnswerUsecase
}

func (e *evaluationUsecase) GetByProgramUser(program_id string, user_id string, stage string) ([]model.Evaluation, error) {
	return e.repo.FilterByProgramUser(program_id, user_id, stage)
}
func (e *evaluationUsecase) GetEvaluateeByProgramPanelist(program_id string, panelist_id string) ([]model.Evaluation, error) {
	return e.repo.FilterByProgramPanelist(program_id, panelist_id)
}

// DeleteData implements EvaluationUsecase
func (e *evaluationUsecase) DeleteData(id string) error {
	// Delete all QuestionAnswer
	qa, err := e.questionAnswer.GetByEvaluation(id)
	if err != nil {
		return err
	}
	for _, q := range qa {
		e.questionAnswer.DeleteData(q.ID)
	}
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
	participant, err := e.participant.FindById(payload.ParticipantID)
	if err != nil {
		return err
	}
	for _, v := range panelist.Roles {
		if v.Name == "judges" {
			payload.Participant = *participant
			return e.repo.Save(payload)
		}
	}
	return fmt.Errorf("Judge assigned is not a valid judge")
}

func NewEvaluationUsecase(repo repository.EvaluationRepo, user UserUsecase, participant ParticipantUsecase, qa QuestionAnswerUsecase) EvaluationUsecase {
	return &evaluationUsecase{
		repo:           repo,
		user:           user,
		participant:    participant,
		questionAnswer: qa,
	}
}
