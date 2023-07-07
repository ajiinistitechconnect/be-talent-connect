package response

import (
	"github.com/alwinihza/talent-connect-be/model"
)

type ProgramResponse struct {
	Program  model.Program
	Activity []Activity
}

type Activity struct {
	Date       string
	Activities []model.Activity
}

type MentoringSchedule struct {
	Date               string
	MentoringSchedules []model.MentoringSchedule
}
