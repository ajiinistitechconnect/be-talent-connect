package repository

import (
	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type QuestionAnswerRepo interface {
	Save(payload *model.QuestionAnswer) error
	Get(id string) (*model.QuestionAnswer, error)
	GetByEvaluation(id string) ([]model.QuestionAnswer, error)
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
	err := q.db.First(&payload, "id = ?", id).Error
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

func NewQuestionAnswerRepo(db *gorm.DB) QuestionAnswerRepo {
	return &questionAnswerRepo{
		db: db,
	}
}
