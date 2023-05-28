package model

type QuestionAnswer struct {
	BaseModel
	QuestionID string
	Question
	AnswerID string
	Answer
	EvaluationQuestionID string
	EvaluationQuestion
}
