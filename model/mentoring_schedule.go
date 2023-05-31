package model

import "time"

type MentoringSchedule struct {
	BaseModel
	MentorMentees []MentorMentee `gorm:"many2many:mentor_mentee_schedules"`
	MentoringDate time.Time
}
