package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuestionCategoryRepo interface {
	BaseRepository[model.QuestionCategory]
}

type questionCategoryRepo struct {
	db *gorm.DB
}

// Delete implements QuestionCategoryRepo
func (q *questionCategoryRepo) Delete(id string) error {
	result := q.db.Delete(&model.QuestionCategory{
		BaseModel: model.BaseModel{
			ID: id,
		},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Question Category not found!")
	}
	return nil
}

// Get implements QuestionCategoryRepo
func (q *questionCategoryRepo) Get(id string) (*model.QuestionCategory, error) {
	var payload model.QuestionCategory
	err := q.db.Preload(clause.Associations).First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

// List implements QuestionCategoryRepo
func (q *questionCategoryRepo) List() ([]model.QuestionCategory, error) {
	var payloads []model.QuestionCategory
	err := q.db.Find(&payloads).Error
	if err != nil {
		return nil, err
	}
	return payloads, nil
}

// Save implements QuestionCategoryRepo
func (q *questionCategoryRepo) Save(payload *model.QuestionCategory) error {
	err := q.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func NewQuestionCategoryRepo(db *gorm.DB) QuestionCategoryRepo {
	return &questionCategoryRepo{
		db: db,
	}
}
