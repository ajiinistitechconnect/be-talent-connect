package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type QuestionCategoryUsecase interface {
	BaseUsecase[model.QuestionCategory]
}

type questionCategoryUsecase struct {
	repo repository.QuestionCategoryRepo
}

// DeleteData implements QuestionCategoryUsecase
func (q *questionCategoryUsecase) DeleteData(id string) error {
	return q.repo.Delete(id)
}

// FindAll implements QuestionCategoryUsecase
func (q *questionCategoryUsecase) FindAll() ([]model.QuestionCategory, error) {
	return q.repo.List()
}

// FindById implements QuestionCategoryUsecase
func (q *questionCategoryUsecase) FindById(id string) (*model.QuestionCategory, error) {
	return q.repo.Get(id)
}

// SaveData implements QuestionCategoryUsecase
func (q *questionCategoryUsecase) SaveData(payload *model.QuestionCategory) error {
	return q.repo.Save(payload)
}

func NewQuestionCategoryUsecase(repo repository.QuestionCategoryRepo) QuestionCategoryUsecase {
	return &questionCategoryUsecase{
		repo: repo,
	}
}
