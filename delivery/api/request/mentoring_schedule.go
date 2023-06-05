package request

import "time"

type MentoringScheduleRequest struct {
	ID            string    `json:"id"`
	MentorMentees []string  `json:"mentorMentees"`
	MentoringDate time.Time `json:"mentoringDate"`
}
