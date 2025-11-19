package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

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
	repo     repository.IQuestionRepository
	hintRepo repository.IHintRepository
	hg       HintGenerator
}

func NewHintsUsecase(repo repository.IQuestionRepository, hintRepo repository.IHintRepository, hg HintGenerator) IHintsUsecase {
	return &hintsUsecase{
		repo:     repo,
		hintRepo: hintRepo,
		hg:       hg,
	}
}

func (u *hintsUsecase) GetHints(ctx context.Context, in GenerateHintInput) (*GenerateHintOutput, error) {

	q, err := u.repo.GetQuestionByID(in.QuestionID)
	if err != nil {
		return nil, fmt.Errorf("問題取得失敗: %w", err)
	}

	// 事前生成されたヒントを検索
	stateKey := generateStateKey(in.Answers)
	hint, err := u.hintRepo.GetHintByQuestionIDAndState(in.QuestionID, stateKey)
	if err == nil {
		// 事前生成されたヒントが見つかった
		var hints []string
		if err := json.Unmarshal([]byte(hint.Hints), &hints); err != nil {
			return nil, fmt.Errorf("ヒントのパース失敗: %w", err)
		}
		return &GenerateHintOutput{
			Correct: false,
			Hints:   hints,
		}, nil
	}

	// 事前生成されたヒントがない場合はリアルタイム生成
	hintList, err := u.hg.Generate(ctx, q, in.Answers)
	if err != nil {
		return nil, fmt.Errorf("ヒント生成失敗: %w", err)
	}

	return &GenerateHintOutput{
		Correct: false,
		Hints:   hintList,
	}, nil
}

func generateStateKey(answers map[string]*string) string {
	var keys []string
	for k, v := range answers {
		// nil でない、かつ空文字列でないキーのみを含める
		if v != nil && *v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return strings.Join(keys, ",")
}
