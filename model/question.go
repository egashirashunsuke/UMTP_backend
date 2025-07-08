package model

import (
	"gorm.io/gorm"
)

type Question struct {
	ID                 int      `json:"id" gorm:"primaryKey"`
	ProblemDescription string   `json:"problem_description"`
	Question           string   `json:"question"`
	Answer             string   `json:"answer"`
	Image              string   `json:"image"`
	Choices            []Choice `json:"choices" gorm:"foreignKey:QuestionID"`
	CreatedAt          string   `json:"created_at"`
}

func GetQuestionByID(db *gorm.DB, id int) (*Question, error) {
	var question Question
	if err := db.Preload("Choices").First(&question, id).Error; err != nil {
		return nil, err
	}
	return &question, nil
}
