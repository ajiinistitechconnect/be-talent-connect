package model

type EvaluationQuestion struct {
	BaseModel
	Evaluation
	EvaluationID     string
	CategoryID       string
	QuestionCategory QuestionCategory
	CategoryWeight   float64
}
