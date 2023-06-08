package model

import "time"

type Activity struct {
	BaseModel
	ProgramID   string  `json:"programID" binding:"required"`
	Program     Program `json:"program,omitempty"`
	Name        string
	Link        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
}
