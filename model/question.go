package model

type Question struct {
	BaseModel
	Question   string
	CategoryId string
}
