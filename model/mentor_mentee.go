package model

type MentorMentee struct {
	BaseModel
	ProgramID     string
	MentorID      string
	Mentor        User
	ParticipantID string
	Participant   User
}
