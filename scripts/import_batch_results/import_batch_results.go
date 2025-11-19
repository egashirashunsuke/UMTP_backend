package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"gorm.io/gorm"
)

type HintsResponse struct {
	Hints []string `json:"hints"`
}

func main() {
	batchID := flag.String("batch-id", "", "OpenAI batch ID")
	flag.Parse()

	if *batchID == "" {
		log.Fatal("-batch-id は必須です")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	db := model.DBConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	client := openai.NewClient(
		option.WithAPIKey(mustGetEnv("OPENAI_API_KEY")),
	)

	if err := importBatchResult(ctx, db, client, *batchID); err != nil {
		log.Fatalf("インポート失敗: %v", err)
	}

	log.Println("✅ インポート完了")
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s が設定されていません", key)
	}
	return v
}

func importBatchResult(ctx context.Context, db *gorm.DB, client openai.Client, batchID string) error {
	// 1. バッチの状態取得
	batch, err := client.Batches.Get(ctx, batchID)
	if err != nil {
		return fmt.Errorf("バッチ取得失敗: %w", err)
	}

	if batch.Status != openai.BatchStatusCompleted {
		return fmt.Errorf("バッチ status=%s のため処理できません", batch.Status)
	}

	if batch.OutputFileID == "" {
		return fmt.Errorf("output_file_id が空です")
	}

	// 2. 出力ファイルを取得
	resp, err := client.Files.Content(ctx, batch.OutputFileID)
	if err != nil {
		return fmt.Errorf("結果ファイル取得失敗: %w", err)
	}
	defer resp.Body.Close()

	return readAndStoreHints(db, resp.Body)
}

type BatchOutputLine struct {
	CustomID string `json:"custom_id"`
	Response struct {
		StatusCode int `json:"status_code"`
		Body       struct {
			Choices []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			} `json:"choices"`
		} `json:"body"`
	} `json:"response"`
	Error *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func readAndStoreHints(db *gorm.DB, r io.Reader) error {
	dec := json.NewDecoder(r)

	for {
		var line BatchOutputLine
		if err := dec.Decode(&line); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("JSONL decode エラー: %w", err)
		}

		if line.Error != nil {
			log.Printf("custom_id=%s はエラー: code=%s msg=%s",
				line.CustomID, line.Error.Code, line.Error.Message)
			continue
		}
		if len(line.Response.Body.Choices) == 0 {
			log.Printf("custom_id=%s choices が空", line.CustomID)
			continue
		}

		content := line.Response.Body.Choices[0].Message.Content

		// content は "{ \"hints\": [ ... ] }" という JSON 文字列のはず
		var hintsResp HintsResponse
		if err := json.Unmarshal([]byte(content), &hintsResp); err != nil {
			log.Printf("custom_id=%s hints JSON parse error: %v", line.CustomID, err)
			continue
		}

		qID, stateKey, err := parseCustomID(line.CustomID)
		if err != nil {
			log.Printf("custom_id=%s parse error: %v", line.CustomID, err)
			continue
		}

		if err := storeHint(db, qID, stateKey, hintsResp.Hints); err != nil {
			log.Printf("custom_id=%s DB 保存エラー: %v", line.CustomID, err)
			continue
		}
	}

	return nil
}

func parseCustomID(customID string) (uint, string, error) {
	parts := strings.SplitN(customID, "_", 2)
	if len(parts) != 2 || !strings.HasPrefix(parts[0], "q") {
		return 0, "", fmt.Errorf("unexpected custom_id format")
	}
	idStr := strings.TrimPrefix(parts[0], "q")

	n, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, "", fmt.Errorf("questionID parse error: %w", err)
	}
	stateKey := parts[1] // "a,b,c" みたいなやつ

	return uint(n), stateKey, nil
}

func storeHint(db *gorm.DB, questionID uint, stateKey string, hints []string) error {
	hintsJSON, err := json.Marshal(hints)
	if err != nil {
		return fmt.Errorf("hints marshal error: %w", err)
	}

	hint := model.Hint{
		QuestionID:   questionID,
		AnswersState: stateKey,
		Hints:        string(hintsJSON),
	}

	return db.
		Where("question_id = ? AND answers_state = ?", questionID, stateKey).
		FirstOrCreate(&hint).Error
}
