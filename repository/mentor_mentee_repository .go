package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MentorMenteeRepo interface {
	BaseRepository[model.MentorMentee]
}

type mentorMenteeRepo struct {
	db *gorm.DB
}

func (m *mentorMenteeRepo) Save(payload *model.MentorMentee) error {
	err := m.db.Save(&payload)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (m *mentorMenteeRepo) Get(id string) (*model.MentorMentee, error) {
	var mentorMentee model.MentorMentee
	err := m.db.Preload(clause.Associations).First(&mentorMentee, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &mentorMentee, nil
}

func (m *mentorMenteeRepo) List() ([]model.MentorMentee, error) {
	var mentorMentees []model.MentorMentee
	err := m.db.Preload(clause.Associations).Find(&mentorMentees).Error
	if err != nil {
		return nil, err
	}
	return mentorMentees, nil
}

func (m *mentorMenteeRepo) Delete(id string) error {
	result := m.db.Delete(&model.MentorMentee{
		BaseModel: model.BaseModel{ID: id},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("MentorMentee not found!")
	}
	return nil
}

func NewMentorMenteeRepo(db *gorm.DB) MentorMenteeRepo {
	return &mentorMenteeRepo{
		db: db,
	}
}
