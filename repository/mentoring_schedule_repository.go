package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type MentoringScheduleRepo interface {
	BaseRepository[model.MentoringSchedule]
	RemoveAllMentorMentees(mentoringScheduleID string) error
	FindByMentorId(mentorId string) ([]model.MentoringSchedule, error)
	FindByMenteeId(mentorId string) ([]model.MentoringSchedule, error)
	SaveFeedback(request *model.MentorMenteeSchedule) error
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

func (m *mentoringScheduleRepo) FindByMentorId(mentorId string) ([]model.MentoringSchedule, error) {
	var mentoringSchedules []model.MentoringSchedule
	err := m.db.Preload("MentorMentees").
		Preload("MentorMentees.Mentor").
		Preload("MentorMentees.Participant").
		Joins("JOIN mentor_mentee_schedules on mentor_mentee_schedules.mentoring_schedule_id = mentoring_schedules.id").
		Joins("JOIN mentor_mentees ON mentor_mentees.id = mentor_mentee_schedules.mentor_mentee_id").
		Where("mentor_mentees.mentor_id = ?", mentorId).
		Find(&mentoringSchedules).Error
	if err != nil {
		return nil, err
	}
	return mentoringSchedules, nil
}

func (m *mentoringScheduleRepo) FindByMenteeId(menteeId string) ([]model.MentoringSchedule, error) {
	var mentoringSchedules []model.MentoringSchedule
	err := m.db.Preload("MentorMentees").
		Preload("MentorMentees.Mentor").
		Preload("MentorMentees.Participant").
		Joins("JOIN mentor_mentee_schedules on mentor_mentee_schedules.mentoring_schedule_id = mentoring_schedules.id").
		Joins("JOIN mentor_mentees ON mentor_mentees.id = mentor_mentee_schedules.mentor_mentee_id").
		Where("mentor_mentees.participant_id = ?", menteeId).
		Find(&mentoringSchedules).Error
	if err != nil {
		return nil, err
	}
	return mentoringSchedules, nil
}

func (m *mentoringScheduleRepo) Get(id string) (*model.MentoringSchedule, error) {
	var mentoringSchedule model.MentoringSchedule
	err := m.db.Preload("MentorMentees").
		Preload("MentorMentees.Mentor").
		Preload("MentorMentees.Participant").
		First(&mentoringSchedule, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	var mentorMenteeSchedules []model.MentorMenteeSchedule
	err = m.db.
		Table("mentor_mentee_schedules").
		Select("mentor_mentee_schedules.*").
		Joins("JOIN mentor_mentees ON mentor_mentees.id = mentor_mentee_schedules.mentor_mentee_id").
		Joins("JOIN mentoring_schedules ON mentoring_schedules.id = mentor_mentee_schedules.mentoring_schedule_id").
		Where("mentoring_schedules.id = ?", id).
		Scan(&mentorMenteeSchedules).Error
	if err != nil {
		return nil, err
	}

	mentoringSchedule.MentorMenteeSchedules = mentorMenteeSchedules
	return &mentoringSchedule, nil
}

func (m *mentoringScheduleRepo) SaveFeedback(request *model.MentorMenteeSchedule) error {
	err := m.db.Save(&request).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *mentoringScheduleRepo) List() ([]model.MentoringSchedule, error) {
	var mentoringSchedules []model.MentoringSchedule
	err := m.db.Preload("MentorMentees").
		Preload("MentorMentees.Mentor").
		Preload("MentorMentees.Participant").
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
