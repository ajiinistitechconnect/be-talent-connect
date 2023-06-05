package model

type Participant struct {
	BaseModel
	ProgramID       string
	UserID          string
	User            User
	EvaluationScore float64
}
