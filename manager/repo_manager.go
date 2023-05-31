package manager

import "github.com/alwinihza/talent-connect-be/repository"

type RepoManager interface {
	RoleRepo() repository.RoleRepo
	UserRepo() repository.UserRepo
	MentoringScheduleRepo() repository.MentoringScheduleRepo
	MentorMenteeRepo() repository.MentorMenteeRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) RoleRepo() repository.RoleRepo {
	return repository.NewRoleRepo(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepo {
	return repository.NewUserRepo(r.infra.Conn())
}

func (r *repoManager) MentoringScheduleRepo() repository.MentoringScheduleRepo {
	return repository.NewMentoringScheduleRepo(r.infra.Conn())
}

func (r *repoManager) MentorMenteeRepo() repository.MentorMenteeRepo {
	return repository.NewMentorMenteeRepo(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
