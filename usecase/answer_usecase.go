package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type AnswerUsecase interface {
	SaveData(payload model.Answer) error
}

type answerUsecase struct {
	repo repository.AnswerRepo
}

// SaveData implements AnswerUsecase.
func (a *answerUsecase) SaveData(payload model.Answer) error {
	return a.repo.Save(&payload)
}

func NewAnswerUsecase(repo repository.AnswerRepo) AnswerUsecase {
	return &answerUsecase{
		repo: repo,
	}
}
