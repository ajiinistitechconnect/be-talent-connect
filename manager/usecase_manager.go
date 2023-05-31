package manager

import (
	"github.com/alwinihza/talent-connect-be/usecase"
)

type UsecaseManager interface {
	UserUc() usecase.UserUsecase
	RoleUc() usecase.RoleUsecase
	MentoringScheduleUc() usecase.MentoringScheduleUsecase
	MentorMenteeUc() usecase.MentorMenteeUsecase
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

func (u *usecaseManager) MentoringScheduleUc() usecase.MentoringScheduleUsecase {
	return usecase.NewMentoringScheduleUsecase(u.repo.MentoringScheduleRepo())
}

func (u *usecaseManager) MentorMenteeUc() usecase.MentorMenteeUsecase {
	return usecase.NewMentorMenteeUsecase(u.repo.MentorMenteeRepo())
}

func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{
		repo: repo,
	}
}
