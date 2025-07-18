package model

type Label struct {
	ID             int             `json:"id" gorm:"primaryKey"`
	QuestionID     int             `json:"question_id"`
	Question       Question        `json:"question" gorm:"foreignKey:QuestionID"`
	LabelCode      string          `json:"label_code"`
	OrderId        int             `json:"order_id"`
	AnswerMappings []AnswerMapping `json:"answer_mappings" gorm:"foreignKey:LabelID"`
}
