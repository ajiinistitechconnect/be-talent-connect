package model

import "github.com/lib/pq"

type Question struct {
	BaseModel
	Question string
	Type     string
	Option   pq.StringArray `gorm:"type:text[]"`
}
