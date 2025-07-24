package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/usecase"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type PromptData struct {
	ProblemDescription   string
	Question             string
	ClassDiagramPlantUML string
	Choices              string
	StudentAnswers       string
	Steps                string
}
type HintsResponse struct {
	Hints []string `json:"hints"`
}

// hintsServiceImpl は実際の実装
type hintsServiceImpl struct {
	client openai.Client
	model  string
}

// NewHintsService は HintsService を生成するコンストラクタ
func NewHintsService() usecase.HintGenerator {
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
func (s *hintsServiceImpl) Generate(ctx context.Context, question *model.Question, answers map[string]*string) ([]string, error) {

	var keys []string
	for k := range answers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for _, k := range keys {
		val := ""
		if answers[k] != nil {
			val = *answers[k]
		}
		sb.WriteString(fmt.Sprintf("%s: %s\n", k, val))
	}
	answersStr := sb.String()

	prompt, err := BuildPrompt(PromptData{
		ProblemDescription:   question.ProblemDescription,
		Question:             question.Question,
		ClassDiagramPlantUML: question.ClassDiagramPlantUML,
		Choices:              FormatChoices(question.Choices),
		StudentAnswers:       answersStr,
		Steps:                question.AnswerProcess,
	})
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("OpenAI からの応答が空です")
	}

	content := resp.Choices[0].Message.Content

	var hintsResp HintsResponse
	if err := json.Unmarshal([]byte(content), &hintsResp); err != nil {
		return nil, fmt.Errorf("failed to parse hints JSON: %w", err)
	}

	return hintsResp.Hints, nil
}

func BuildPrompt(data PromptData) (string, error) {
	tmpl, err := template.ParseFiles("template/hints_prompt.tmpl")
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func FormatChoices(choices []model.Choice) string {
	// ラベル順にソート
	sort.Slice(choices, func(i, j int) bool {
		return choices[i].ChoiceCode < choices[j].ChoiceCode
	})
	var sb strings.Builder
	for _, c := range choices {
		sb.WriteString(fmt.Sprintf("%s. %s\n", c.ChoiceCode, c.ChoiceText))
	}
	return sb.String()
}
