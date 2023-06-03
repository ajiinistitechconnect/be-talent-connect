package model

type Participant struct {
	BaseModel
	ProgramID       string
	Program         Program
	UserID          string
	User            User
	EvaluationScore float64
}
