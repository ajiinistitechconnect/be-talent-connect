package model

type Evaluation struct {
	BaseModel
	PanelistID    string
	Panelist      User
	ParticipantID string
	Participant   Participant
	Score         float64 `gorm:"default:0.0"`
}
