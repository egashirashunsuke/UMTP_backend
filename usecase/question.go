package usecase

import (
	"fmt"

	"github.com/egashirashunsuke/UMTP_backend/dto"
	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
)

type IQuestionUsecase interface {
	GetQuestionByID(id int) (model.Question, error)
	GetAllQuestions() ([]model.Question, error)
	CreateQuestion(in *dto.CreateQuestionDTO) error
	GetNextQuestion(currentID int) (model.Question, error)
	GetPrevQuestion(currentID int) (model.Question, error)
	CheckAnswer(questionID int, answers map[string]*string) (bool, error)
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

func (uc *questionUsecase) GetAllQuestions() ([]model.Question, error) {
	q, err := uc.qr.GetAllQuestions()
	if err != nil {
		return []model.Question{}, fmt.Errorf("failed to get question by ID %d: %w", err)
	}
	return *q, nil
}

func (uc *questionUsecase) CreateQuestion(in *dto.CreateQuestionDTO) error {
	if err := uc.qr.CreateQuestionWithAssociations(in); err != nil {
		return fmt.Errorf("failed to create question: %w", err)
	}
	return nil
}

func (uc *questionUsecase) GetNextQuestion(currentID int) (model.Question, error) {
	// 現在のIDより大きい最小のIDを持つ質問を取得
	nextQuestion, err := uc.qr.GetNextQuestionByID(currentID + 1)
	if err != nil {
		return model.Question{}, fmt.Errorf("failed to get next question after ID %d: %w", currentID, err)
	}
	return nextQuestion, nil
}

func (uc *questionUsecase) GetPrevQuestion(currentID int) (model.Question, error) {
	// 現在のIDより小さい最大のIDを持つ質問を取得
	prevQuestion, err := uc.qr.GetPrevQuestionByID(currentID - 1)
	if err != nil {
		return model.Question{}, fmt.Errorf("failed to get previous question before ID %d: %w", currentID, err)
	}
	return prevQuestion, nil
}

func (uc *questionUsecase) CheckAnswer(questionID int, answers map[string]*string) (bool, error) {
	q, err := uc.qr.GetQuestionByID(questionID)
	if err != nil {
		return false, fmt.Errorf("問題取得失敗: %w", err)
	}

	ok, msg := q.Check(answers)
	fmt.Println(msg)
	return ok, nil
}
