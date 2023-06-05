package model

type Question struct {
	BaseModel
	Question    string
	Type        string
	Description string
	Options     []Option `gorm:"foreignKey:question_id" json:"options,omitempty"`
}

type Option struct {
	ID         string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	QuestionID string
	Value      int
	Text       string
}
