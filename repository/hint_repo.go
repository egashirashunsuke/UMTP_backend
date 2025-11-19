package repository

import (
	"github.com/egashirashunsuke/UMTP_backend/model"
	"gorm.io/gorm"
)

type IHintRepository interface {
	GetHintByQuestionIDAndState(questionID int, answersState string) (*model.Hint, error)
}

type hintRepository struct {
	db *gorm.DB
}

func NewHintRepository(db *gorm.DB) IHintRepository {
	return &hintRepository{db: db}
}

func (r *hintRepository) GetHintByQuestionIDAndState(questionID int, answersState string) (*model.Hint, error) {
	var hint model.Hint
	if err := r.db.Where("question_id = ? AND answers_state = ?", questionID, answersState).
		First(&hint).Error; err != nil {
		return nil, err
	}
	return &hint, nil
}
