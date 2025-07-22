package service

import (
	"context"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"gorm.io/gorm"
)

type QuestionService interface {
	CheckAnswer(ctx context.Context, question *model.Question, answers map[string]*string) (bool, error)
}

type questionServiceImpl struct {
	DB *gorm.DB
}

func NewQuestionService(db *gorm.DB) QuestionService {
	return &questionServiceImpl{DB: db}
}

func (s *questionServiceImpl) CheckAnswer(ctx context.Context, question *model.Question, answers map[string]*string) (bool, error) {
	// 正解マッピングを作成
	correct := make(map[string]string) // label_code -> choice_code
	for _, am := range question.AnswerMappings {
		// ChoiceとLabelをPreloadしておく必要あり
		correct[am.Label.LabelCode] = am.Choice.ChoiceCode
	}

	// 回答を比較
	for label, userChoicePtr := range answers {
		userChoice := ""
		if userChoicePtr != nil {
			userChoice = *userChoicePtr
		}
		if correctChoice, ok := correct[label]; !ok || userChoice != correctChoice {
			return false, nil // 不正解
		}
	}

	return true, nil // 全て一致なら正解
}
