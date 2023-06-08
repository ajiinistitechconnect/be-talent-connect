package model

import "time"

type MentoringSchedule struct {
	BaseModel
	MentorMentees []MentorMentee `gorm:"many2many:mentor_mentee_schedules" json:"mentorMentees,omitempty"`
	Name          string
	Link          string
	Description   string
	StartDate     time.Time
	EndDate       time.Time
}
