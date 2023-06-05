package model

type Evaluation struct {
	BaseModel
	PanelistID      string `json:"panelistId"`
	Panelist        User
	ParticipantID   string `json:"participantId"`
	Participant     Participant
	QuestionAnswers []QuestionAnswer `gorm:"many2many:evaluation_answers"`
	Stage           string
	IsEvaluated     bool    `gorm:"not null,default:false"`
	Score           float64 `gorm:"default:0.0"`
}
