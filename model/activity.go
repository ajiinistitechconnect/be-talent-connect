package model

import "time"

type Activity struct {
	BaseModel
	Program
	ProgramID string
	Name      string
	StartDate time.Time
	Duration  time.Duration
}
