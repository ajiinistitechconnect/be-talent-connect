package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ParticipantRepo interface {
	BaseRepository[model.Participant]
	GetByProgram(id string) ([]model.Participant, error)
	GetEvaluationByStage(participant_id string, stage string) (model.Participant, error)
}

type participantRepo struct {
	db *gorm.DB
}

func (p *participantRepo) GetEvaluationByStage(participant_id string, stage string) (model.Participant, error) {
	var payload model.Participant
	err := p.db.Preload("Evaluations", "evaluations.stage = ?", stage).First(&payload, "participants.id = ?", participant_id).Error
	if err != nil {
		return model.Participant{}, err
	}
	return payload, nil
}

func (m *participantRepo) Save(payload *model.Participant) error {
	err := m.db.Save(&payload)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (m *participantRepo) Get(id string) (*model.Participant, error) {
	var participant model.Participant
	err := m.db.Preload(clause.Associations).First(&participant, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &participant, nil
}

func (m *participantRepo) List() ([]model.Participant, error) {
	var participants []model.Participant
	err := m.db.Preload(clause.Associations).Find(&participants).Error
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (m *participantRepo) GetByProgram(id string) ([]model.Participant, error) {
	var participants []model.Participant
	err := m.db.Preload(clause.Associations).Find(&participants, "program_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (m *participantRepo) Delete(id string) error {
	result := m.db.Delete(&model.Participant{
		BaseModel: model.BaseModel{ID: id},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Participant not found!")
	}
	return nil
}

func NewParticipantRepo(db *gorm.DB) ParticipantRepo {
	return &participantRepo{
		db: db,
	}
}
