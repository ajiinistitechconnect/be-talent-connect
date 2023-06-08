package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type QuestionRepo interface {
	BaseRepository[model.Question]
	Update(payload *model.Question) error
}

type questionRepo struct {
	db *gorm.DB
}

// Delete implements QuestionRepo
func (q *questionRepo) Delete(id string) error {
	result := q.db.Delete(&model.Question{
		BaseModel: model.BaseModel{
			ID: id,
		},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Program not found!")
	}
	return nil
}

// Get implements QuestionRepo
func (q *questionRepo) Get(id string) (*model.Question, error) {
	var payload model.Question
	err := q.db.Preload("Options", func(db *gorm.DB) *gorm.DB {
		return db.Order("options.value")
	}).First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

// List implements QuestionRepo
func (q *questionRepo) List() ([]model.Question, error) {
	var payloads []model.Question
	err := q.db.Preload("Options", func(db *gorm.DB) *gorm.DB {
		return db.Order("options.value")
	}).Find(&payloads).Error
	if err != nil {
		return nil, err
	}
	return payloads, nil
}

func (q *questionRepo) Update(payload *model.Question) error {
	err := q.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

// Save implements QuestionRepo
func (q *questionRepo) Save(payload *model.Question) error {
	err := q.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func NewQuestionRepo(db *gorm.DB) QuestionRepo {
	return &questionRepo{
		db: db,
	}
}
