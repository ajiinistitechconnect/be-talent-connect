package manager

import (
	"github.com/alwinihza/talent-connect-be/usecase"
)

type UsecaseManager interface {
	UserUc() usecase.UserUsecase
	RoleUc() usecase.RoleUsecase
	MentoringScheduleUsecase() usecase.MentoringScheduleUsecase
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

func (u *usecaseManager) MentoringScheduleUsecase() usecase.MentoringScheduleUsecase {
	return usecase.NewMentoringScheduleUsecase(u.repo.MentoringScheduleRepo())
}

func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{
		repo: repo,
	}
}
