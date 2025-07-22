package usecase

import (
	"fmt"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
)

type IQuestionUsecase interface {
	GetQuestionByID(id int) (model.Question, error)
}

type questionUsecase struct {
	qr repository.IQuestionRepository
}

func NewQuestionUsecase(qr repository.IQuestionRepository) IQuestionUsecase {
	return &questionUsecase{qr: qr}
}

func (uc *questionUsecase) GetQuestionByID(id int) (model.Question, error) {
	q, err := uc.qr.GetQuestionByID(id)
	if err != nil {
		return model.Question{}, fmt.Errorf("failed to get question by ID %d: %w", id, err)
	}
	return *q, nil
}
