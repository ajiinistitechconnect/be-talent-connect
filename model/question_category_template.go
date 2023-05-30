package model

type QuestionCategory struct {
	BaseModel
	Name        string
	Description string
	Questions   []Question `gorm:"foreignKey:category_id"`
}
