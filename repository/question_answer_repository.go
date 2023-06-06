package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type QuestionAnswerRepo interface {
	Save(payload *model.QuestionAnswer) error
	Get(id string) (*model.QuestionAnswer, error)
	GetByEvaluation(id string) ([]model.QuestionAnswer, error)
	Delete(id string) error
	GetByQuestion(evaluation_id string, category_id string, question_id string) (*model.QuestionAnswer, error)
}

type questionAnswerRepo struct {
	db *gorm.DB
}

// GetByEvaluation implements QuestionAnswerRepo.
func (q *questionAnswerRepo) GetByEvaluation(id string) ([]model.QuestionAnswer, error) {
	var payload []model.QuestionAnswer
	err := q.db.Find(&payload, "evaluation_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// Get implements QuestionAnswerRepo.
func (q *questionAnswerRepo) Get(id string) (*model.QuestionAnswer, error) {
	var payload model.QuestionAnswer
	err := q.db.Preload("Answer").First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

// Save implements QuestionAnswerRepo.
func (q *questionAnswerRepo) Save(payload *model.QuestionAnswer) error {
	err := q.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (q *questionAnswerRepo) Delete(id string) error {
	result := q.db.Delete(&model.QuestionAnswer{
		BaseModel: model.BaseModel{
			ID: id,
		},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Question Answer not found!")
	}
	return nil
}

func (q *questionAnswerRepo) GetByQuestion(evaluation_id string, category_id string, question_id string) (*model.QuestionAnswer, error) {
	var payload model.QuestionAnswer
	err := q.db.Preload("Answer").First(&payload, "evaluation_id = ? and evaluation_category_question_id = ? and question_id = ? ", evaluation_id, category_id, question_id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func NewQuestionAnswerRepo(db *gorm.DB) QuestionAnswerRepo {
	return &questionAnswerRepo{
		db: db,
	}
}
