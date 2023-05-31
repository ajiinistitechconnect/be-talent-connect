package model

import "time"

type MentoringSchedule struct {
	BaseModel
	MentorMentees []MentorMentee `gorm:"many2many:mentor_mentee_schedules" json:"mentorMentees,omitempty"`
	MentoringDate time.Time      `json:"mentoringDate,omitempty"`
}
