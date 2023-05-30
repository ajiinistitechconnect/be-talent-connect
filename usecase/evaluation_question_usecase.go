package usecase

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type EvaluationQuestionUsecase interface {
	BaseUsecase[model.EvaluationQuestion]
}

type evaluationQuestionUsecase struct {
	repo       repository.EvaluationQuestionRepo
	evaluation EvaluationUsecase
}

// DeleteData implements EvaluationQuestionUsecase
func (e *evaluationQuestionUsecase) DeleteData(id string) error {
	return e.repo.Delete(id)
}

// FindAll implements EvaluationQuestionUsecase
func (e *evaluationQuestionUsecase) FindAll() ([]model.EvaluationQuestion, error) {
	return e.repo.List()
}

// FindById implements EvaluationQuestionUsecase
func (e *evaluationQuestionUsecase) FindById(id string) (*model.EvaluationQuestion, error) {
	return e.repo.Get(id)
}

// SaveData implements EvaluationQuestionUsecase
func (e *evaluationQuestionUsecase) SaveData(payload *model.EvaluationQuestion) error {
	evaluation, err := e.evaluation.FindById(payload.EvaluationID)
	if err != nil {
		return err
	}
	if payload.CategoryWeight < 0 {
		return fmt.Errorf("Weight cannot less than 0")
	}
	if total, err := e.evaluation.GetTotalWeight(payload.EvaluationID); err != nil {
		return err
	} else if total+payload.CategoryWeight > 100.0 {
		return fmt.Errorf("Weight total exceed 100")
	}
	payload.Evaluation = *evaluation
	return e.repo.Save(payload)
}

func NewEvaluationQuestionUsecase(repo repository.EvaluationQuestionRepo) EvaluationQuestionUsecase {
	return &evaluationQuestionUsecase{
		repo: repo,
	}
}
