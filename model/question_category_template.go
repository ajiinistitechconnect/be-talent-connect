package model

type QuestionCategoryTemplate struct {
	BaseModel
	Name        string
	Description string
	Questions   []Question `gorm:"foreignKey:category_id"`
}
