package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type TotalWeight struct {
	id    string
	total float64
}

type ProgramRepo interface {
	BaseRepository[model.Program]
	BaseSearch[model.Program]
}

type programRepo struct {
	db *gorm.DB
}

func (p *programRepo) Save(payload *model.Program) error {
	err := p.db.Save(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *programRepo) Get(id string) (*model.Program, error) {
	var payload model.Program
	err := p.db.First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (p *programRepo) List() ([]model.Program, error) {
	var programs []model.Program
	err := p.db.Find(&programs).Error
	if err != nil {
		return nil, err
	}
	return programs, nil
}

func (p *programRepo) Delete(id string) error {
	result := p.db.Delete(&model.Program{
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

func (p *programRepo) Search(by map[string]any) ([]model.Program, error) {
	var programs []model.Program
	result := p.db.Where(by).Find(&programs)
	if err := result.Error; err != nil {
		return nil, err
	}
	return programs, nil
}

func NewProgramRepo(db *gorm.DB) ProgramRepo {
	return &programRepo{
		db: db,
	}
}
