package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // React の開発サーバー
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		// 必要に応じて PUT, DELETE, OPTIONS なども追加
	}))
	e.GET("/", func(c echo.Context) error {
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
	})

	e.Logger.Fatal(e.Start(":8000"))
}
