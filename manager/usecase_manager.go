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
	ParticipantUc() usecase.ParticipantUsecase
	AuthUc() usecase.AuthUsecase
	QuestionUc() usecase.QuestionUsecase
	QuestionCategoryUc() usecase.QuestionCategoryUsecase
	EvaluationCategoryUc() usecase.EvaluationCategoryUsecase
	EvaluationUc() usecase.EvaluationUsecase
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

func (u *usecaseManager) ParticipantUc() usecase.ParticipantUsecase {
	return usecase.NewParticipantUsecase(u.repo.ParticipantRepo(), u.UserUc(), u.ProgramUc())
}

func (u *usecaseManager) AuthUc() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(u.UserUc())
}

func (u *usecaseManager) QuestionUc() usecase.QuestionUsecase {
	return usecase.NewQuestionUsecase(u.repo.QuestionRepo())
}

func (u *usecaseManager) QuestionCategoryUc() usecase.QuestionCategoryUsecase {
	return usecase.NewQuestionCategoryUsecase(u.repo.QuestionCategoryRepo(), u.QuestionUc())
}

func (u *usecaseManager) EvaluationUc() usecase.EvaluationUsecase {
	return usecase.NewEvaluationUsecase(u.repo.EvaluationRepo(), u.UserUc(), u.ParticipantUc())
}

func (u *usecaseManager) EvaluationCategoryUc() usecase.EvaluationCategoryUsecase {
	return usecase.NewEvaluationQuestionUsecase(u.repo.EvaluationCategoryRepo(), u.EvaluationUc(), u.ProgramUc(), u.QuestionCategoryUc())
}

func NewUsecaseManager(repo RepoManager, cfg *config.Config) UsecaseManager {
	return &usecaseManager{
		repo: repo,
		cfg:  cfg,
	}
}
