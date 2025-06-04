package handler

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func CreateHintsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(".env ファイルの読み込みに失敗しました")
		}

		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("APIキーが環境変数に設定されていません")
		}
		client := openai.NewClient(
			option.WithAPIKey(apiKey),
		)
		chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("日本語に対応していますか？"),
			},
			Model: openai.ChatModelGPT4o,
		})
		if err != nil {
			panic(err.Error())
		}
		println(chatCompletion.Choices[0].Message.Content)

		return c.String(http.StatusOK, chatCompletion.Choices[0].Message.Content)
	}
}
