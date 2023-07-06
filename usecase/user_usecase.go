package usecase

import (
	"github.com/alwinihza/talent-connect-be/config"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
	"github.com/alwinihza/talent-connect-be/utils"
	"gorm.io/gorm"
)

type UserUsecase interface {
	BaseUsecase[model.User]
	SearchEmail(email string) (*model.User, error)
	UpdateRole(payload *model.User, role []string) error
	UpdateData(payload *model.User) error
	SearchByRole(role string) ([]model.User, error)
	SearchAvailableMenteeForProgram(program_id string, name string) ([]model.User, error)
	SearchAvailableMenteeForMentor(mentor_id string, program_id string, name string) ([]model.User, error)
	SearchAvailableMenteeForJudges(panelist_id_id string, program_id string, name string) ([]model.User, error)
}

type userUsecase struct {
	repo repository.UserRepo
	role RoleUsecase
	cfg  *config.Config
}

func (u *userUsecase) SearchAvailableMenteeForProgram(program_id string, name string) ([]model.User, error) {
	return u.repo.SearchForMenteeProgram(program_id, name)
}

func (u *userUsecase) SearchAvailableMenteeForMentor(mentor_id string, program_id string, name string) ([]model.User, error) {
	return u.repo.SearchForMentee(program_id, mentor_id, name)
}

func (u *userUsecase) SearchAvailableMenteeForJudges(panelist_id string, program_id string, name string) ([]model.User, error) {
	return u.repo.SearchMenteeForJudges(program_id, panelist_id, name)
}

func (u *userUsecase) FindAll() ([]model.User, error) {
	return u.repo.List()
}

func (u *userUsecase) FindById(id string) (*model.User, error) {
	return u.repo.Get(id)
}

func (u *userUsecase) SearchEmail(email string) (*model.User, error) {
	return u.repo.SearchByEmail(email)
}

func (u *userUsecase) SaveData(payload *model.User) error {
	_, err := u.SearchEmail(payload.Email)
	if err != gorm.ErrRecordNotFound {
		return err
	}
	// password, err := utils.GeneratePassword()
	// if err != nil {
	// 	return err
	// }
	payload.Password, err = utils.SaltPassword([]byte("password"))
	if err != nil {
		return err
	}
	if err := u.repo.Save(payload); err != nil {
		return err
	}
	// body := fmt.Sprintf("Hi %s, You are registered to TalentConnect Platform\n\nYour Password is <b>%s</b>", payload.FirstName, password)
	// log.Println(body)
	// if err := utils.SendMail([]string{payload.Email}, "TalentConnect Registration", body, u.cfg.SMTPConfig); err != nil {
	// 	return err
	// }
	return nil
}

func (u *userUsecase) UpdateData(payload *model.User) error {
	return u.repo.Update(payload)
}

func (u *userUsecase) UpdateRole(payload *model.User, role []string) error {
	for _, v := range role {
		tempRole, err := u.role.FindByName(v)
		if err != nil {
			return err
		}
		payload.Roles = append(payload.Roles, *tempRole)
	}
	return nil
}

func (u *userUsecase) DeleteData(id string) error {
	return u.repo.Delete(id)
}

func (u *userUsecase) SearchByRole(role string) ([]model.User, error) {
	return u.repo.SearchByRole(role)
}

func NewUserUseCase(repo repository.UserRepo, role RoleUsecase, cfg *config.Config) UserUsecase {
	return &userUsecase{
		repo: repo,
		role: role,
		cfg:  cfg,
	}
}
