package request

import "time"

type MentoringScheduleRequest struct {
	ID            string
	MentorMentees []string
	Name          string
	Link          string
	Description   string
	StartDate     time.Time
	EndDate       time.Time
}
