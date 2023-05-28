package model

type EvaluationQuestion struct {
	BaseModel
	Evaluation
	EvaluationID   string
	CategoryWeight int
}
