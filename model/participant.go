package model

import "database/sql"

type Participant struct {
	BaseModel
	ProgramID            string
	UserID               string
	User                 User
	MidEvaluationScore   sql.NullFloat64
	FinalEvaluationScore sql.NullFloat64
	Evaluations          []Evaluation
}
