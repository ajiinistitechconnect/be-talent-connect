package model

import "time"

type Activity struct {
	BaseModel
	ProgramID string `json:"programID" binding:"required"`
	Name      string
	StartDate time.Time
}
