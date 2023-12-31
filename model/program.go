package model

import "time"

type Program struct {
	BaseModel
	Name                 string
	Participants         []Participant                `gorm:"foreignKey:program_id" json:"participants,omitempty"`
	Mentors              []MentorMentee               `gorm:"foreignKey:program_id" json:"mentors,omitempty"`
	Activities           []Activity                   `gorm:"foreignKey:program_id" json:"activities,omitempty"`
	EvaluationCategories []EvaluationCategoryQuestion `gorm:"foreignKey:program_id"`
	StartDate            time.Time                    `json:"startDate"`
	EndDate              time.Time                    `json:"endDate"`
}
