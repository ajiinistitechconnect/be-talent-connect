package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type EvaluationCategoryRepo interface {
	BaseRepository[model.EvaluationCategoryQuestion]
	AggregateWeight(id string) (float64, error)
	GetQuestions(program_id string) ([]model.EvaluationCategoryQuestion, error)
}

type evaluationQuestionRepo struct {
	db *gorm.DB
}

// Delete implements EvaluationQuestionRepo
func (e *evaluationQuestionRepo) Delete(id string) error {
	result := e.db.Delete(&model.EvaluationCategoryQuestion{
		BaseModel: model.BaseModel{
			ID: id,
		},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Evaluation Category not found!")
	}
	return nil
}

// Get implements EvaluationQuestionRepo
func (e *evaluationQuestionRepo) Get(id string) (*model.EvaluationCategoryQuestion, error) {
	var payload model.EvaluationCategoryQuestion
	err := e.db.First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

// List implements EvaluationQuestionRepo
func (e *evaluationQuestionRepo) List() ([]model.EvaluationCategoryQuestion, error) {
	var payloads []model.EvaluationCategoryQuestion
	err := e.db.Find(&payloads).Error
	if err != nil {
		return nil, err
	}
	return payloads, nil
}

// Save implements EvaluationQuestionRepo
func (e *evaluationQuestionRepo) Save(payload *model.EvaluationCategoryQuestion) error {
	err := e.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *evaluationQuestionRepo) AggregateWeight(id string) (float64, error) {
	var ret TotalWeight
	result := e.db.Model(model.EvaluationCategoryQuestion{}).Select("sum(category_weight) as total").Where("program_id = ?", id).First(&ret)
	if result.Error != nil {
		return 0.0, result.Error
	}
	return ret.total, nil
}

func (e *evaluationQuestionRepo) GetQuestions(program_id string) ([]model.EvaluationCategoryQuestion, error) {
	var ret []model.EvaluationCategoryQuestion
	err := e.db.Preload("QuestionCategory").Preload("QuestionCategory.Questions").Find(&ret, "program_id = ?", program_id).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func NewEvaluationQuestionRepo(db *gorm.DB) EvaluationCategoryRepo {
	return &evaluationQuestionRepo{
		db: db,
	}
}
