package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// HintsService はビジネスロジック層のインターフェース（テスト用に実装差し替えしやすくするため）
type HintsService interface {
	GetHints(ctx context.Context, answer map[string]*string) (string, error)
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
func (s *hintsServiceImpl) GetHints(ctx context.Context, answers map[string]*string) (string, error) {

	templatePath := filepath.Join("template", "hints_prompt.tmpl")
	b, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	var keys []string
	for k := range answers {
		keys = append(keys, k)
	}
	var sb strings.Builder
	for _, k := range keys {
		val := ""
		if answers[k] != nil {
			val = *answers[k]
		}
		sb.WriteString(fmt.Sprintf("%s: %s\n", k, val))
	}
	answersStr := sb.String()

	prompt := fmt.Sprintf(string(b), answersStr)

	fmt.Println("=== OpenAIへ送るプロンプト ===")
	fmt.Println(prompt)

	// ChatCompletion のリクエスト用にメッセージを作る
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.UserMessage(prompt),
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
