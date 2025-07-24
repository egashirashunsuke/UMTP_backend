package usecase

import (
	"context"
	"fmt"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
)

type GenerateHintInput struct {
	QuestionID int
	Answers    map[string]*string
}
type GenerateHintOutput struct {
	Correct bool     `json:"correct"`
	Message string   `json:"message"`
	Hints   []string `json:"hints"` // ヒントのリスト
}

type IHintsUsecase interface {
	GetHints(ctx context.Context, in GenerateHintInput) (*GenerateHintOutput, error)
}

type HintGenerator interface {
	Generate(ctx context.Context, q *model.Question, answers map[string]*string) ([]string, error)
}

type hintsUsecase struct {
	repo repository.IQuestionRepository
	hg   HintGenerator
}

func NewHintsUsecase(repo repository.IQuestionRepository, hg HintGenerator) IHintsUsecase {
	return &hintsUsecase{repo: repo, hg: hg}
}

func (u *hintsUsecase) GetHints(ctx context.Context, in GenerateHintInput) (*GenerateHintOutput, error) {

	q, err := u.repo.GetQuestionByID(in.QuestionID)
	if err != nil {
		return nil, fmt.Errorf("問題取得失敗: %w", err)
	}

	// 2) 正誤判定（モデル側にメソッドがある想定）
	ok, msg := q.Check(in.Answers) // 例: (bool, string) を返す
	if ok {
		return &GenerateHintOutput{
			Correct: true,
			Message: msg,
		}, nil
	}

	// 3) 不正解ならヒント生成（AI など）
	hint, err := u.hg.Generate(ctx, q, in.Answers)
	if err != nil {
		return nil, fmt.Errorf("ヒント生成失敗: %w", err)
	}

	return &GenerateHintOutput{
		Correct: false,
		Hints:   hint,
	}, nil
}
