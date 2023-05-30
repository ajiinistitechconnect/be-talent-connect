package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type ActivityRepo interface {
	BaseRepository[model.Activity]
}

type activityRepo struct {
	db *gorm.DB
}

// Delete implements ActivityRepo
func (a *activityRepo) Delete(id string) error {
	result := a.db.Delete(&model.Program{
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

// Get implements ActivityRepo
func (a *activityRepo) Get(id string) (*model.Activity, error) {
	var payload model.Activity
	err := a.db.First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

// List implements ActivityRepo
func (a *activityRepo) List() ([]model.Activity, error) {
	var payloads []model.Activity
	err := a.db.Find(&payloads).Error
	if err != nil {
		return nil, err
	}
	return payloads, nil
}

// Save implements ActivityRepo
func (a *activityRepo) Save(payload *model.Activity) error {
	err := a.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func NewActivityRepo(db *gorm.DB) ActivityRepo {
	return &activityRepo{
		db: db,
	}
}
