package model

type Hint struct {
	ID           int      `json:"id" gorm:"primaryKey"`
	QuestionID   uint     `json:"question_id"`
	Question     Question `json:"question" gorm:"foreignKey:QuestionID"`
	AnswersState string   `json:"answers_state" gorm:"uniqueIndex:idx_question_state"`
	Hints        string   `json:"hints" gorm:"type:text"` // JSON配列として保存
	CreatedAt    int64    `gorm:"autoCreateTime"`
}
