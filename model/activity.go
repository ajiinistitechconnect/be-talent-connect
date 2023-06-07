package model

import "time"

type Activity struct {
	BaseModel
	ProgramID   string `json:"programID" binding:"required"`
	Name        string
	Link        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
}
