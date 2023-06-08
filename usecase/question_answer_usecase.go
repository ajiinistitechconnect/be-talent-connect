package usecase

import (
	"strconv"

	"github.com/alwinihza/talent-connect-be/delivery/api/request"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type QuestionAnswerUsecase interface {
	SaveQuestionAnswer(payload *request.AnswerRequest) error
	SaveData(payload *model.QuestionAnswer) error
	FindById(id string) (*model.QuestionAnswer, error)
	DeleteData(id string) error
	GetByEvaluation(id string) ([]model.QuestionAnswer, error)
	ScoreByCategory(evaluation_id string, category_id string) (float64, error)
	UpdateEvaluationScore(evaluation_id string, program_id string) error
}

type questionAnswerUsecase struct {
	repo                   repository.QuestionAnswerRepo
	answer                 AnswerUsecase
	evaluationRepo         repository.EvaluationRepo
	evaluationCategoryRepo repository.EvaluationCategoryRepo
	program                ProgramUsecase
	question               QuestionUsecase
}

// Get implements QuestionAnswerUsecase.
func (q *questionAnswerUsecase) FindById(id string) (*model.QuestionAnswer, error) {
	payload, err := q.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// GetByEvaluation implements QuestionAnswerUsecase.
func (q *questionAnswerUsecase) GetByEvaluation(id string) ([]model.QuestionAnswer, error) {
	return q.repo.GetByEvaluation(id)
}

func (q *questionAnswerUsecase) SaveQuestionAnswer(payload *request.AnswerRequest) error {
	for _, QuestionCategories := range payload.QuestionCategories {
		for _, v := range QuestionCategories.QuestionList {
			temp := model.QuestionAnswer{
				QuestionID:                   v.QuestionID,
				EvaluationID:                 payload.EvaluationID,
				EvaluationCategoryQuestionID: QuestionCategories.CategoryID,
				Answer: model.Answer{
					Answer: v.Answer,
				},
			}
			if err := q.SaveData(&temp); err != nil {
				return err
			}
		}
	}
	if err := q.UpdateEvaluationScore(payload.EvaluationID, payload.ProgramID); err != nil {
		return err
	}
	return nil
}

// SaveData implements QuestionAnswerUsecase.
func (q *questionAnswerUsecase) SaveData(payload *model.QuestionAnswer) error {
	if _, err := q.evaluationRepo.Get(payload.EvaluationID); err != nil {
		return err
	}

	if _, err := q.evaluationCategoryRepo.Get(payload.EvaluationCategoryQuestionID); err != nil {
		return err
	}

	return q.repo.Save(payload)
}

func (q *questionAnswerUsecase) DeleteData(id string) error {
	return q.repo.Delete(id)
}

// ScoreByCategory implements QuestionAnswerUsecase.
func (q *questionAnswerUsecase) ScoreByCategory(evaluation_id string, category_id string) (float64, error) {
	question_category, err := q.evaluationCategoryRepo.Get(category_id)
	if err != nil {
		return 0.0, err
	}
	var question_id []string
	for _, v := range question_category.QuestionCategory.Questions {
		if v.Type == "rating" {
			question_id = append(question_id, v.ID)
		}
	}
	var score float64
	for _, v := range question_id {
		question, err := q.question.FindById(v)
		if err != nil {
			return 0.0, err
		}
		answer, err := q.repo.GetByQuestion(evaluation_id, category_id, v)
		if err != nil {
			continue
		}
		ans, err := strconv.Atoi(answer.Answer.Answer)
		if err != nil {
			return 0.0, err
		}
		max_option := len(question.Options)
		score += float64(ans) / float64(max_option) * 100
	}
	score /= float64(len(question_id))
	return score, nil
}

// UpdateEvaluationScore implements QuestionAnswerUsecase.
// Use it in controller answer
func (q *questionAnswerUsecase) UpdateEvaluationScore(evaluation_id string, program_id string) error {
	// retrieve all EvaluationCategory
	categoryList, err := q.program.ListQuestions(program_id)
	if err != nil {
		return err
	}
	var finalScore float64
	for _, v := range categoryList.EvaluationCategories {
		evaluationCategory, err := q.evaluationCategoryRepo.Get(v.ID)
		if err != nil {
			return err
		}
		score, err := q.ScoreByCategory(evaluation_id, v.ID)
		if err != nil {
			return err
		}
		tempScore := score * (evaluationCategory.CategoryWeight / 100.0)
		finalScore += tempScore
	}
	evaluation, err := q.evaluationRepo.Get(evaluation_id)
	if err != nil {
		return err
	}
	evaluation.IsEvaluated = true
	evaluation.Score = finalScore
	if err := q.evaluationRepo.Save(evaluation); err != nil {
		return err
	}
	return nil
}

func NewQuestionAnswerUsecase(
	repo repository.QuestionAnswerRepo,
	answer AnswerUsecase,
	evaluationRepo repository.EvaluationRepo,
	evaluationCat repository.EvaluationCategoryRepo,
	program ProgramUsecase,
	question QuestionUsecase,
) QuestionAnswerUsecase {
	return &questionAnswerUsecase{
		repo:                   repo,
		answer:                 answer,
		evaluationRepo:         evaluationRepo,
		evaluationCategoryRepo: evaluationCat,
		program:                program,
		question:               question,
	}
}
