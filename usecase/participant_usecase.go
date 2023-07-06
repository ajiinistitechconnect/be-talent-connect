package usecase

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type ParticipantUsecase interface {
	BaseUsecase[model.Participant]
	GetEvaluationScore(participant_id string) (model.EvaluationResult, error)
	ListEvaluationByStage(participant_id string, stage string) (model.Participant, error)
	CheckEvaluationScore(participant_id string) (model.EvaluationResult, error)
	ListByProgram(program_id string) ([]model.Participant, error)
}

type participantUsecase struct {
	repo    repository.ParticipantRepo
	user    UserUsecase
	program ProgramUsecase
}

func (p *participantUsecase) ListByProgram(program_id string) ([]model.Participant, error) {
	return p.repo.GetByProgram(program_id)
}

func (p *participantUsecase) ListEvaluationByStage(participant_id string, stage string) (model.Participant, error) {
	return p.repo.GetEvaluationByStage(participant_id, stage)
}

func (p *participantUsecase) GetEvaluationScore(participant_id string) (model.EvaluationResult, error) {
	result, err := p.CheckEvaluationScore(participant_id)
	if err != nil {
		return model.EvaluationResult{}, nil
	}
	return result, nil
}

func (p *participantUsecase) CheckEvaluationScore(participant_id string) (model.EvaluationResult, error) {
	var finalResult model.EvaluationResult
	mid_finished := true
	mid_count := 0.0
	participant, err := p.ListEvaluationByStage(participant_id, "mid")
	if err != nil {
		return model.EvaluationResult{}, err
	}
	if !participant.MidEvaluationScore.Valid {
		for _, pa := range participant.Evaluations {
			if !pa.IsEvaluated {
				mid_finished = false
				continue
			}
			if mid_finished {
				finalResult.MidScore += pa.Score
				mid_count++
			}
		}
		if !mid_finished {
			finalResult.MidScore = 0.0
			finalResult.MidStatus = "Not Finished"
		} else {
			finalResult.MidScore /= mid_count
			if finalResult.MidScore > 70.0 {
				finalResult.MidStatus = "Passed"
			} else {
				finalResult.MidStatus = "Failed"
			}
		}
	} else {
		finalResult.MidScore = participant.MidEvaluationScore.Float64
		if finalResult.MidScore > 70.0 {
			finalResult.MidStatus = "Passed"
		} else {
			finalResult.MidStatus = "Failed"
		}
	}

	participant, err = p.ListEvaluationByStage(participant_id, "final")
	if err != nil {
		return model.EvaluationResult{}, err
	}

	final_finished := true
	final_count := 0.0
	if !participant.FinalEvaluationScore.Valid {
		for _, pa := range participant.Evaluations {
			if !pa.IsEvaluated {
				final_finished = false
				continue
			}
			if final_finished {
				finalResult.FinalScore += pa.Score
				final_count++
			}
		}

		if !final_finished {
			finalResult.FinalScore = 0.0
			finalResult.FinalStatus = "Not Finished"
		} else {
			finalResult.FinalScore /= final_count
			if finalResult.FinalScore > 70.0 {
				finalResult.FinalStatus = "Passed"
			} else {
				finalResult.FinalStatus = "Failed"
			}
		}
	} else {
		finalResult.FinalScore = participant.FinalEvaluationScore.Float64
		if finalResult.FinalScore > 70.0 {
			finalResult.FinalStatus = "Passed"
		} else {
			finalResult.FinalStatus = "Failed"
		}
	}
	return finalResult, nil
}

func (m *participantUsecase) FindAll() ([]model.Participant, error) {
	return m.repo.List()
}

func (m *participantUsecase) FindById(id string) (*model.Participant, error) {
	return m.repo.Get(id)
}

func (m *participantUsecase) SaveData(payload *model.Participant) error {
	_, err := m.program.FindById(payload.ProgramID)
	if err != nil {
		return err
	}

	user, err := m.user.FindById(payload.UserID)
	if err != nil {
		return err
	}
	for _, v := range user.Roles {
		if v.Name == "mentee" {
			payload.User = *user

			return m.repo.Save(payload)
		}
	}
	return fmt.Errorf("Participant assigned is not a valid participant")
}

func (m *participantUsecase) DeleteData(id string) error {
	return m.repo.Delete(id)
}

func NewParticipantUsecase(
	repo repository.ParticipantRepo,
	user UserUsecase,
	program ProgramUsecase,
) ParticipantUsecase {
	return &participantUsecase{
		repo:    repo,
		user:    user,
		program: program,
	}
}
