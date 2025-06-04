package service

import (
	"context"
	"errors"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// HintsService はビジネスロジック層のインターフェース（テスト用に実装差し替えしやすくするため）
type HintsService interface {
	GetHints(ctx context.Context, question string) (string, error)
}

// hintsServiceImpl は実際の実装
type hintsServiceImpl struct {
	client openai.Client
	model  string
}

// NewHintsService は HintsService を生成するコンストラクタ
func NewHintsService() HintsService {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		// アプリ起動時に .env または環境変数に設定がないとパニックにする
		panic("OPENAI_API_KEY が設定されていません")
	}
	// OpenAI クライアントを初期化
	cli := openai.NewClient(option.WithAPIKey(apiKey))

	// モデル名は環境変数 OPENAI_MODEL で上書きできるようにする
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = string(openai.ChatModelGPT4o) // デフォルトを GPT-4o に設定
	}

	return &hintsServiceImpl{
		client: cli,
		model:  model,
	}
}

// GetHints は OpenAI の ChatCompletion を呼び出し、返ってきたテキストを返す
func (s *hintsServiceImpl) GetHints(ctx context.Context, question string) (string, error) {
	if question == "" {
		return "", errors.New("質問が空です")
	}

	// ChatCompletion のリクエスト用にメッセージを作る
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.UserMessage(question),
	}

	// ChatCompletion を実行
	resp, err := s.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    s.model,
	})
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("OpenAI からの応答が空です")
	}

	// 返ってきた最初の選択肢を文字列として返却
	return resp.Choices[0].Message.Content, nil
}
