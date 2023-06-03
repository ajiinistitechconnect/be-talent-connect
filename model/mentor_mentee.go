package model

type MentorMentee struct {
	BaseModel
	ProgramID     string
	Program       Program
	MentorID      string
	Mentor        User
	ParticipantID string
	Participant   User
}
