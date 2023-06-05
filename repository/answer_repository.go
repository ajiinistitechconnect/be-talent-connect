package repository

import (
	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type AnswerRepo interface {
	Save(payload *model.Answer) error
}

type answerRepo struct {
	db *gorm.DB
}

// Save implements AnswerRepo.
func (a *answerRepo) Save(payload *model.Answer) error {
	err := a.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func NewAnswerRepo(db *gorm.DB) AnswerRepo {
	return &answerRepo{
		db: db,
	}
}
