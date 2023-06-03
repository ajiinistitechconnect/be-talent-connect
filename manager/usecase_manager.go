package manager

import (
	"github.com/alwinihza/talent-connect-be/config"
	"github.com/alwinihza/talent-connect-be/usecase"
)

type UsecaseManager interface {
	UserUc() usecase.UserUsecase
	RoleUc() usecase.RoleUsecase
	MentoringScheduleUc() usecase.MentoringScheduleUsecase
	MentorMenteeUc() usecase.MentorMenteeUsecase
	ProgramUc() usecase.ProgramUsecase
	ActivityUc() usecase.ActivityUsecase
	AuthUc() usecase.AuthUsecase
}

type usecaseManager struct {
	repo RepoManager
	cfg  *config.Config
}

func (u *usecaseManager) RoleUc() usecase.RoleUsecase {
	return usecase.NewRoleUsecase(u.repo.RoleRepo())
}

func (u *usecaseManager) UserUc() usecase.UserUsecase {
	return usecase.NewUserUseCase(u.repo.UserRepo(), u.RoleUc(), u.cfg)
}

func (u *usecaseManager) MentoringScheduleUc() usecase.MentoringScheduleUsecase {
	return usecase.NewMentoringScheduleUsecase(u.repo.MentoringScheduleRepo(), u.MentorMenteeUc())
}

func (u *usecaseManager) MentorMenteeUc() usecase.MentorMenteeUsecase {
	return usecase.NewMentorMenteeUsecase(u.repo.MentorMenteeRepo(), u.UserUc(), u.ProgramUc())
}

func (u *usecaseManager) ProgramUc() usecase.ProgramUsecase {
	return usecase.NewProgramUsecase(u.repo.ProgramRepo())
}

func (u *usecaseManager) ActivityUc() usecase.ActivityUsecase {
	return usecase.NewActivityUsecase(u.repo.ActivityRepo(), u.ProgramUc())
}

func (u *usecaseManager) AuthUc() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(u.UserUc())
}

func NewUsecaseManager(repo RepoManager, cfg *config.Config) UsecaseManager {
	return &usecaseManager{
		repo: repo,
		cfg:  cfg,
	}
}
