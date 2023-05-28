package model

import "time"

type MentoringSchedule struct {
	BaseModel
	MentorMenteeID string
	MentorMentee
	MentoringDate time.Time
}
