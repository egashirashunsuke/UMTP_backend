package model

type AnswerMapping struct {
	ID         int      `json:"id" gorm:"primaryKey"`
	QuestionID int      `json:"question_id"`
	Question   Question `json:"question" gorm:"foreignKey:QuestionID"`
	ChoiceID   int      `json:"choice_id"`
	Choice     Choice   `json:"choice" gorm:"foreignKey:ChoiceID"`
	LabelID    int      `json:"label_id"`
	Label      Label    `json:"label" gorm:"foreignKey:LabelID"`
}
