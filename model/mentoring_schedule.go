package model

import (
	"time"

	"gorm.io/gorm"
)

type MentoringSchedule struct {
	BaseModel
	MentorMentees         []MentorMentee `gorm:"many2many:mentor_mentee_schedules" json:"mentorMentees,omitempty"`
	MentorMenteeSchedules []MentorMenteeSchedule
	Name                  string
	Link                  string
	Description           string
	StartDate             time.Time
	EndDate               time.Time
}

type MentorMenteeSchedule struct {
	CreatedAt           time.Time      `gorm:"<-:create" json:"-"`
	UpdatedAt           time.Time      `json:"-"`
	DeletedAt           gorm.DeletedAt `json:"-"`
	MentoringScheduleID string         `gorm:"primaryKey;type:uuid"`
	MentorMenteeID      string         `gorm:"primaryKey;type:uuid"`
	Date                time.Time
	Comment             string
}
