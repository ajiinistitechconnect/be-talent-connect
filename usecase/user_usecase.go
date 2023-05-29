package usecase

import (
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type UserUsecase interface {
	BaseUsecase[model.User]
	UpdateRole(payload *model.User, role []string) error
	UpdateData(payload *model.User) error
}

type userUsecase struct {
	repo repository.UserRepo
	role RoleUsecase
}

func (u *userUsecase) FindAll() ([]model.User, error) {
	return u.repo.List()
}

func (u *userUsecase) FindById(id string) (*model.User, error) {
	return u.repo.Get(id)
}

func (u *userUsecase) SaveData(payload *model.User) error {
	return u.repo.Save(payload)
}

func (u *userUsecase) UpdateData(payload *model.User) error {
	return u.repo.Update(payload)
}

func (u *userUsecase) UpdateRole(payload *model.User, role []string) error {
	for _, v := range role {
		tempRole, err := u.role.FindByName(v)
		if err != nil && tempRole != nil {
			return err
		}
		payload.Roles = append(payload.Roles, *tempRole)
	}
	return nil
}

func (u *userUsecase) DeleteData(id string) error {
	return u.repo.Delete(id)
}

func NewUserUseCase(repo repository.UserRepo, role RoleUsecase) UserUsecase {
	return &userUsecase{
		repo: repo,
		role: role,
	}
}
