package model

type Evaluation struct {
	BaseModel
	PanelistID    string
	Panelist      User
	ParticipantID string
	// Participant     Participant
	QuestionAnswers []QuestionAnswer `gorm:"many2many:evaluation_answers"`
	Stage           string
	IsEvaluated     bool    `gorm:"default:false"`
	Score           float64 `gorm:"default:0.0"`
}
