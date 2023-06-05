package model

type EvaluationCategoryQuestion struct {
	BaseModel
	ProgramID          string
	CategoryWeight     float64
	QuestionCategoryID string
	QuestionCategory   QuestionCategory
}
