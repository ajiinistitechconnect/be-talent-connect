package model

type Participant struct {
	BaseModel
	ProgramID string
	Program
	UserID string
	User
	EvaluationScore float64
}
