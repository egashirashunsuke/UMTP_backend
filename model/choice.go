package model

type Choice struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	QuestionID int    `json:"question_id"`
	Label      string `json:"label"`
	Text       string `json:"text"`
}
