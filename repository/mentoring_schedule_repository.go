package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type MentoringScheduleRepo interface {
	BaseRepository[model.MentoringSchedule]
	RemoveAllMentorMentees(mentoringScheduleID string) error
}

type mentoringScheduleRepo struct {
	db *gorm.DB
}

func (m *mentoringScheduleRepo) Save(payload *model.MentoringSchedule) error {
	err := m.db.Save(&payload)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (m *mentoringScheduleRepo) Get(id string) (*model.MentoringSchedule, error) {
	var mentoringSchedule model.MentoringSchedule
	err := m.db.Preload("MentorMentees").
		Preload("MentorMentees.Mentor").
		Preload("MentorMentees.Participant").
		Preload("MentorMentees.Program").
		First(&mentoringSchedule, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &mentoringSchedule, nil
}

func (m *mentoringScheduleRepo) List() ([]model.MentoringSchedule, error) {
	var mentoringSchedules []model.MentoringSchedule
	err := m.db.Preload("MentorMentees").
		Preload("MentorMentees.Mentor").
		Preload("MentorMentees.Participant").
		Preload("MentorMentees.Program").
		Find(&mentoringSchedules).Error
	if err != nil {
		return nil, err
	}
	return mentoringSchedules, nil
}

func (m *mentoringScheduleRepo) RemoveAllMentorMentees(mentoringScheduleID string) error {
	mentoringSchedule := &model.MentoringSchedule{
		BaseModel: model.BaseModel{
			ID: mentoringScheduleID,
		},
	}

	err := m.db.Model(mentoringSchedule).Association("MentorMentees").Clear()
	if err != nil {
		return err
	}

	return nil
}

func (m *mentoringScheduleRepo) Delete(id string) error {
	result := m.db.Delete(&model.MentoringSchedule{
		BaseModel: model.BaseModel{ID: id},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("MentoringSchedule not found!")
	}
	return nil
}

func NewMentoringScheduleRepo(db *gorm.DB) MentoringScheduleRepo {
	return &mentoringScheduleRepo{
		db: db,
	}
}
