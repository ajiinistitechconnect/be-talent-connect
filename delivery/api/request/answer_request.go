package request

type AnswerRequest struct {
	EvaluationID       string
	PanelistID         string
	QuestionCategories []QuestionCategory
}

type QuestionCategory struct {
	CategoryID   string
	QuestionList []QuestionList
}

type QuestionList struct {
	QuestionID string
	Answer     string
}
