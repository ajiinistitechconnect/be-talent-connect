package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type EvaluationRepo interface {
	BaseRepository[model.Evaluation]
}

type evaluationRepo struct {
	db *gorm.DB
}

// Delete implements EvaluationRepo
func (e *evaluationRepo) Delete(id string) error {
	result := e.db.Delete(&model.Evaluation{
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

// Get implements EvaluationRepo
func (e *evaluationRepo) Get(id string) (*model.Evaluation, error) {
	var payload model.Evaluation
	err := e.db.First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

// List implements EvaluationRepo
func (e *evaluationRepo) List() ([]model.Evaluation, error) {
	var payloads []model.Evaluation
	err := e.db.Find(&payloads).Error
	if err != nil {
		return nil, err
	}
	return payloads, nil
}

// Save implements EvaluationRepo
func (e *evaluationRepo) Save(payload *model.Evaluation) error {
	err := e.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func NewEvaluationRepo(db *gorm.DB) EvaluationRepo {
	return &evaluationRepo{
		db: db,
	}
}
