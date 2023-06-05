package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type QuestionAnswerUsecase interface {
	SaveData(payload *model.QuestionAnswer) error
	Get(id string) (*model.QuestionAnswer, error)
	GetByEvaluation(id string) ([]model.QuestionAnswer, error)
	ScoreByCategory(evaluation_id string, category_id string) (float64, error)
	UpdateEvaluationScore(evaluation_id string) error
}

type questionAnswerUsecase struct {
	repo               repository.QuestionAnswerRepo
	answer             AnswerUsecase
	evaluation         EvaluationUsecase
	evaluationCategory EvaluationCategoryUsecase
}

// Get implements QuestionAnswerUsecase.
func (q *questionAnswerUsecase) Get(id string) (*model.QuestionAnswer, error) {
	return q.repo.Get(id)
}

// GetByEvaluation implements QuestionAnswerUsecase.
func (q *questionAnswerUsecase) GetByEvaluation(id string) ([]model.QuestionAnswer, error) {
	return q.repo.GetByEvaluation(id)
}

// SaveData implements QuestionAnswerUsecase.
func (q *questionAnswerUsecase) SaveData(payload *model.QuestionAnswer) error {
	if _, err := q.evaluation.FindById(payload.EvaluationID); err != nil {
		return err
	}

	if _, err := q.evaluationCategory.FindById(payload.EvaluationCategoryQuestionID); err != nil {
		return err
	}

	if err := q.answer.SaveData(payload.Answer); err != nil {
		return err
	}
	payload.AnswerID = payload.Answer.ID

	return q.repo.Save(payload)
}

// ScoreByCategory implements QuestionAnswerUsecase.
func (q *questionAnswerUsecase) ScoreByCategory(evaluation_id string, category_id string) (float64, error) {
	// get all question id that is option
	// count the question id
	// aggregate total answer
	// return the total answer/count
	panic("unimplemented")
}

// UpdateEvaluationScore implements QuestionAnswerUsecase.
func (*questionAnswerUsecase) UpdateEvaluationScore(evaluation_id string) error {
	// retrieve all EvaluationCategory
	// ScoreByCategory * EvaluationWeight
	// Update the Score
	panic("unimplemented")
}

func NewQuestionAnswerUsecase(repo repository.QuestionAnswerRepo, answer AnswerUsecase) QuestionAnswerUsecase {
	return &questionAnswerUsecase{
		repo:   repo,
		answer: answer,
	}
}
