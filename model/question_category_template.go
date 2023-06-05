package model

type QuestionCategory struct {
	BaseModel
	Name        string
	Description string
	Questions   []Question `gorm:"many2many:category_template" json:"questions,omitempty"`
}
