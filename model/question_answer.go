package model

type QuestionAnswer struct {
	BaseModel
	QuestionID                   string
	EvaluationID                 string
	EvaluationCategoryQuestionID string
	AnswerID                     string
	Answer                       Answer
}
