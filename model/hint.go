package model

type Hint struct {
	ID           int      `json:"id" gorm:"primaryKey"`
	QuestionID   uint     `json:"question_id"`
	Question     Question `json:"question" gorm:"foreignKey:QuestionID"`
	AnswersState string   `json:"answers_state"`
	Hints        string   `json:"hints" gorm:"type:text"`
}
