package model

import "time"

type Program struct {
	BaseModel
	Name      string
	StartDate time.Time
	EndDate   time.Time
}
