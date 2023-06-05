package usecase

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type QuestionUsecase interface {
	BaseUsecase[model.Question]
}

type questionUsecase struct {
	repo repository.QuestionRepo
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
	if payload.Type == "text" || payload.Type == "rating" {
		if err := q.repo.Save(payload); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Type not acceptable")
}

func NewQuestionUsecase(repo repository.QuestionRepo) QuestionUsecase {
	return &questionUsecase{
		repo: repo,
	}
}
