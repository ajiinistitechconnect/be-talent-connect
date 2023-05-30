package manager

import "github.com/alwinihza/talent-connect-be/repository"

type RepoManager interface {
	RoleRepo() repository.RoleRepo
	UserRepo() repository.UserRepo
	ProgramRepo() repository.ProgramRepo
	ActivityRepo() repository.ActivityRepo
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

func (r *repoManager) ProgramRepo() repository.ProgramRepo {
	return repository.NewProgramRepo(r.infra.Conn())
}

func (r *repoManager) ActivityRepo() repository.ActivityRepo {
	return repository.NewActivityRepo(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
