package model

type Evaluation struct {
	BaseModel
	PanelistID    string
	Panelist      User
	ParticipantID string
	Participant   User
	Score         int
}
