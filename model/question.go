package model

type Question struct {
	BaseModel
	Question                 string
	CategoryId               string           `json:"categoryId" binding:"required"`
	QuestionCategoryTemplate QuestionCategory `gorm:"foreignKey:CategoryId"`
}
