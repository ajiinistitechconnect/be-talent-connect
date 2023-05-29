package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	BaseRepository[model.User]
	Update(*model.User) error
}

type userRepo struct {
	db *gorm.DB
}

func (u *userRepo) Save(payload *model.User) error {
	err := u.db.Save(&payload)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (u *userRepo) Update(payload *model.User) error {
	if err := u.db.Model(&payload).Association("Roles").Replace(payload.Roles); err != nil {
		return err
	}
	if err := u.db.Save(&payload); err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *userRepo) Get(id string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Roles").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) List() ([]model.User, error) {
	var users []model.User
	err := r.db.Preload("Roles").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Delete(id string) error {
	result := r.db.Delete(&model.User{
		BaseModel: model.BaseModel{ID: id},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Role not found!")
	}
	return nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
