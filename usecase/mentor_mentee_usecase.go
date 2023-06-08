package usecase

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type MentorMenteeUsecase interface {
	BaseUsecase[model.MentorMentee]
	FindByMentorId(id string) ([]model.MentorMentee, error)
	FindByMentorIdProgramId(program_id string, mentor_id string) ([]model.MentorMentee, error)
}

type mentorMenteeUsecase struct {
	repo    repository.MentorMenteeRepo
	user    UserUsecase
	program ProgramUsecase
}

func (m *mentorMenteeUsecase) FindAll() ([]model.MentorMentee, error) {
	return m.repo.List()
}

func (m *mentorMenteeUsecase) FindById(id string) (*model.MentorMentee, error) {
	return m.repo.Get(id)
}

func (m *mentorMenteeUsecase) FindByMentorId(id string) ([]model.MentorMentee, error) {
	return m.repo.GetMentee(id)
}

func (m *mentorMenteeUsecase) FindByMentorIdProgramId(program_id string, mentor_id string) ([]model.MentorMentee, error) {
	return m.repo.GetMenteeProgram(program_id, mentor_id)
}

func (m *mentorMenteeUsecase) SaveData(payload *model.MentorMentee) error {
	_, err := m.program.FindById(payload.ProgramID)
	if err != nil {
		return err
	}
	// payload.Program = *program

	mentor, err := m.user.FindById(payload.MentorID)
	if err != nil {
		return err
	}
	var flag bool
	for _, v := range mentor.Roles {
		if v.Name == "mentor" {
			flag = true
		}
	}
	if !flag {
		return fmt.Errorf("Mentor assigned is not a valid mentor")
	}
	payload.Mentor = *mentor
	flag = false
	mentee, err := m.user.FindById(payload.ParticipantID)
	if err != nil {
		return err
	}
	for _, v := range mentee.Roles {
		if v.Name == "participant" {
			flag = true
		}
	}

	if !flag {
		return fmt.Errorf("Mentee assigned is not a valid mentee")
	}
	payload.Participant = *mentee

	return m.repo.Save(payload)
}

func (m *mentorMenteeUsecase) DeleteData(id string) error {
	return m.repo.Delete(id)
}

func NewMentorMenteeUsecase(
	repo repository.MentorMenteeRepo,
	user UserUsecase,
	program ProgramUsecase,
) MentorMenteeUsecase {
	return &mentorMenteeUsecase{
		repo:    repo,
		user:    user,
		program: program,
	}
}
