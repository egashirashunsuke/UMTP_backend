package model

type Choice struct {
	ID             int             `json:"id" gorm:"primaryKey"`
	QuestionID     int             `json:"question_id"`
	Question       Question        `json:"question" gorm:"foreignKey:QuestionID"`
	ChoiceCode     string          `json:"label"`
	ChoiceText     string          `json:"text"`
	AnswerMappings []AnswerMapping `json:"answer_mappings" gorm:"foreignKey:ChoiceID"`
}
