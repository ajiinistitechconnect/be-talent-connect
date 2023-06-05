package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type QuestionUsecase interface {
	BaseUsecase[model.Question]
}

type questionUsecase struct {
	repo repository.QuestionRepo
	qc   QuestionCategoryUsecase
}

// DeleteData implements QuestionUsecase
func (q *questionUsecase) DeleteData(id string) error {
	return q.repo.Delete(id)
}

// FindAll implements QuestionUsecase
func (q *questionUsecase) FindAll() ([]model.Question, error) {
	return q.repo.List()
}

// FindById implements QuestionUsecase
func (q *questionUsecase) FindById(id string) (*model.Question, error) {
	return q.repo.Get(id)
}

// SaveData implements QuestionUsecase
func (q *questionUsecase) SaveData(payload *model.Question) error {
	return q.repo.Save(payload)
}

func NewQuestionUsecase(repo repository.QuestionRepo, qc QuestionCategoryUsecase) QuestionUsecase {
	return &questionUsecase{
		repo: repo,
		qc:   qc,
	}
}
