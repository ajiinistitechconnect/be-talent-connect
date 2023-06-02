package repository

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	BaseRepository[model.User]
	Update(*model.User) error
	SearchByEmail(email string) (*model.User, error)
	SearchByRole(role string) ([]model.User, error)
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

func (u *userRepo) Get(id string) (*model.User, error) {
	var user model.User
	err := u.db.Preload("Roles").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) List() ([]model.User, error) {
	var users []model.User
	err := u.db.Preload("Roles").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepo) Delete(id string) error {
	result := u.db.Delete(&model.User{
		BaseModel: model.BaseModel{ID: id},
	})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("Uer not found!")
	}
	return nil
}

func (u *userRepo) SearchByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Preload("Roles").First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) SearchByRole(role string) ([]model.User, error) {
	var user []model.User
	err := u.db.Preload("Roles").
		Joins("JOIN users_roles ON users_roles.user_id = users.id").
		Joins("JOIN roles ON users_roles.role_id = roles.id").
		Find(&user, "roles.name = ?", role).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
