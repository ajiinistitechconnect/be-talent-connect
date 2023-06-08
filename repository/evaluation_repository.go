package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EvaluationRepo interface {
	BaseRepository[model.Evaluation]
	FilterByProgramUser(program_id string, user_id string, stage string) ([]model.Evaluation, error)
	FilterByProgramPanelist(program_id string, panelist_id string) ([]model.Evaluation, error)
}

type evaluationRepo struct {
	db *gorm.DB
}

func (e *evaluationRepo) FilterByProgramUser(program_id string, user_id string, stage string) ([]model.Evaluation, error) {
	var payload []model.Evaluation
	err := e.db.Preload("Participant").Preload("Participant.User").Joins("JOIN participants ON participants.id=evaluations.participant_id").
		// Joins("JOIN users on users.id=participants.id").
		Where("evaluations.panelist_id = ? AND participants.program_id = ? AND evaluations.stage = ?", user_id, program_id, stage).
		Find(&payload).Error
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (e *evaluationRepo) FilterByProgramPanelist(program_id string, panelist_id string) ([]model.Evaluation, error) {
	var payload []model.Evaluation
	err := e.db.Preload("Participant").Preload("Participant.User").Joins("JOIN participants ON participants.id=evaluations.participant_id").
		// Joins("JOIN users on users.id=participants.id").
		Where("evaluations.panelist_id = ? AND participants.program_id = ? and evaluations.stage = 'final'", panelist_id, program_id).
		Find(&payload).Error
	if err != nil {
		return nil, err
	}
	return payload, nil

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
	err := e.db.Preload(clause.Associations).First(&payload, "id = ?", id).Error
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
