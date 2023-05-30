package manager

import (
	"github.com/alwinihza/talent-connect-be/usecase"
)

type UsecaseManager interface {
	UserUc() usecase.UserUsecase
	RoleUc() usecase.RoleUsecase
	ProgramUc() usecase.ProgramUsecase
	ActivityUc() usecase.ActivityUsecase
}

type usecaseManager struct {
	repo RepoManager
}

func (u *usecaseManager) RoleUc() usecase.RoleUsecase {
	return usecase.NewRoleUsecase(u.repo.RoleRepo())
}

func (u *usecaseManager) UserUc() usecase.UserUsecase {
	return usecase.NewUserUseCase(u.repo.UserRepo(), u.RoleUc())
}

func (u *usecaseManager) ProgramUc() usecase.ProgramUsecase {
	return usecase.NewProgramUsecase(u.repo.ProgramRepo())
}

func (u *usecaseManager) ActivityUc() usecase.ActivityUsecase {
	return usecase.NewActivityUsecase(u.repo.ActivityRepo(), u.ProgramUc())
}

func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{
		repo: repo,
	}
}
