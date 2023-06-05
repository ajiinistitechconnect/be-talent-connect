package usecase

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type EvaluationCategoryUsecase interface {
	BaseUsecase[model.EvaluationCategoryQuestion]
}

type evaluationCategoryUsecase struct {
	repo       repository.EvaluationCategoryRepo
	evaluation EvaluationUsecase
	program    ProgramUsecase
	category   QuestionCategoryUsecase
}

// DeleteData implements EvaluationQuestionUsecase
func (e *evaluationCategoryUsecase) DeleteData(id string) error {
	return e.repo.Delete(id)
}

// FindAll implements EvaluationQuestionUsecase
func (e *evaluationCategoryUsecase) FindAll() ([]model.EvaluationCategoryQuestion, error) {
	return e.repo.List()
}

// FindById implements EvaluationQuestionUsecase
func (e *evaluationCategoryUsecase) FindById(id string) (*model.EvaluationCategoryQuestion, error) {
	return e.repo.Get(id)
}

// SaveData implements EvaluationQuestionUsecase
func (e *evaluationCategoryUsecase) SaveData(payload *model.EvaluationCategoryQuestion) error {
	if payload.CategoryWeight < 0 {
		return fmt.Errorf("Weight cannot less than 0")
	}

	if _, err := e.program.FindById(payload.ProgramID); err != nil {
		return err
	}

	questionCategory, err := e.category.FindById(payload.QuestionCategoryID)
	if err != nil {
		return err
	}

	payload.QuestionCategory = *questionCategory

	if total, err := e.repo.AggregateWeight(payload.ProgramID); err != nil {
		return err
	} else if total+payload.CategoryWeight > 100.0 {
		return fmt.Errorf("Weight total exceed 100")
	} else {
		fmt.Println(total)
	}
	return e.repo.Save(payload)
}

func NewEvaluationQuestionUsecase(
	repo repository.EvaluationCategoryRepo,
	evaluation EvaluationUsecase,
	program ProgramUsecase,
	category QuestionCategoryUsecase,
) EvaluationCategoryUsecase {
	return &evaluationCategoryUsecase{
		repo:       repo,
		evaluation: evaluation,
		program:    program,
		category:   category,
	}
}
