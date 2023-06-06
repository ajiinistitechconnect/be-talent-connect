package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EvaluationCategoryRepo interface {
	BaseRepository[model.EvaluationCategoryQuestion]
	AggregateWeight(id string) (float64, error)
	GetQuestions(program_id string) ([]model.EvaluationCategoryQuestion, error)
	GetByProgramID(id string) ([]model.EvaluationCategoryQuestion, error)
}

type evaluationCategoryRepo struct {
	db *gorm.DB
}

type TotalWeight struct {
	Id    string
	Total float64
}

// Delete implements EvaluationQuestionRepo
func (e *evaluationCategoryRepo) Delete(id string) error {
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
func (e *evaluationCategoryRepo) Get(id string) (*model.EvaluationCategoryQuestion, error) {
	var payload model.EvaluationCategoryQuestion
	err := e.db.Preload(clause.Associations).Preload("QuestionCategory.Questions").First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (e *evaluationCategoryRepo) GetByProgramID(id string) ([]model.EvaluationCategoryQuestion, error) {
	var payload []model.EvaluationCategoryQuestion
	err := e.db.Preload(clause.Associations).Preload("QuestionCategory.Questions").Find(&payload, "program_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// List implements EvaluationQuestionRepo
func (e *evaluationCategoryRepo) List() ([]model.EvaluationCategoryQuestion, error) {
	var payloads []model.EvaluationCategoryQuestion
	err := e.db.Find(&payloads).Error
	if err != nil {
		return nil, err
	}
	return payloads, nil
}

// Save implements EvaluationQuestionRepo
func (e *evaluationCategoryRepo) Save(payload *model.EvaluationCategoryQuestion) error {
	err := e.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *evaluationCategoryRepo) AggregateWeight(id string) (float64, error) {
	var ret TotalWeight
	result := e.db.Model(&model.EvaluationCategoryQuestion{}).Select("program_id as id, sum(category_weight) as total").Where("program_id = ?", id).Group("program_id").Find(&ret)
	if result.Error != nil {
		return 0.0, result.Error
	}
	return ret.Total, nil
}

func (e *evaluationCategoryRepo) GetQuestions(program_id string) ([]model.EvaluationCategoryQuestion, error) {
	var ret []model.EvaluationCategoryQuestion
	err := e.db.Preload("QuestionCategory").Preload("QuestionCategory.Questions").Find(&ret, "program_id = ?", program_id).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func NewEvaluationCategoryRepo(db *gorm.DB) EvaluationCategoryRepo {
	return &evaluationCategoryRepo{
		db: db,
	}
}
