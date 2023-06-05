package model

type EvaluationCategoryQuestion struct {
	BaseModel
	ProgramID          string `json:"programId"`
	CategoryWeight     float64
	QuestionCategoryID string `json:"questionCategoryId"`
	QuestionCategory   QuestionCategory
}
