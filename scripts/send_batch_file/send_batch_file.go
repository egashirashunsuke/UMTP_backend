package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY が設定されていません")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey))

	f, err := os.Open("batch-input.jsonl")
	if err != nil {
		log.Fatalf("batch-input.jsonl を開けません: %v", err)
	}
	defer f.Close()

	file, err := client.Files.New(ctx, openai.FileNewParams{
		File:    f,
		Purpose: openai.FilePurpose("batch"),
	})
	if err != nil {
		log.Fatalf("ファイルアップロード失敗: %v", err)
	}

	batch, err := client.Batches.New(ctx, openai.BatchNewParams{
		InputFileID:      file.ID,
		Endpoint:         openai.BatchNewParamsEndpoint("/v1/chat/completions"),
		CompletionWindow: openai.BatchNewParamsCompletionWindow24h,
	})
	if err != nil {
		log.Fatalf("バッチ作成失敗: %v", err)
	}
	fmt.Println(batch.ID, batch.Status)

}
