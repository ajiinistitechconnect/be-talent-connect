package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type EvaluationQuestionRepo interface {
	BaseRepository[model.EvaluationQuestion]
}

type evaluationQuestionRepo struct {
	db *gorm.DB
}

// Delete implements EvaluationQuestionRepo
func (e *evaluationQuestionRepo) Delete(id string) error {
	result := e.db.Delete(&model.EvaluationQuestion{
		BaseModel: model.BaseModel{
			ID: id,
		},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Activity not found!")
	}
	return nil
}

// Get implements EvaluationQuestionRepo
func (e *evaluationQuestionRepo) Get(id string) (*model.EvaluationQuestion, error) {
	var payload model.EvaluationQuestion
	err := e.db.First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

// List implements EvaluationQuestionRepo
func (e *evaluationQuestionRepo) List() ([]model.EvaluationQuestion, error) {
	var payloads []model.EvaluationQuestion
	err := e.db.Find(&payloads).Error
	if err != nil {
		return nil, err
	}
	return payloads, nil
}

// Save implements EvaluationQuestionRepo
func (e *evaluationQuestionRepo) Save(payload *model.EvaluationQuestion) error {
	err := e.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func NewEvaluationQuestionRepo(db *gorm.DB) EvaluationQuestionRepo {
	return &evaluationQuestionRepo{
		db: db,
	}
}
