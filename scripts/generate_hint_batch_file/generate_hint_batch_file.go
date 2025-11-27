package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
	"github.com/egashirashunsuke/UMTP_backend/service"
)

type BatchLine struct {
	CustomID string                 `json:"custom_id"`
	Method   string                 `json:"method"`
	URL      string                 `json:"url"`
	Body     map[string]interface{} `json:"body"`
}

func main() {
	targetQuestionID := flag.Uint("question-id", 0, "バッチ生成対象の問題ID (0なら全問題)")
	flag.Parse()

	db := model.DBConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	qRepo := repository.NewQuestionRepository(db)
	questions, err := qRepo.GetAllQuestions()
	if err != nil {
		log.Fatalf("問題取得失敗: %v", err)
	}
	log.Printf("問題数: %d", len(*questions))

	f, err := os.Create("batch-input.jsonl")
	if err != nil {
		log.Fatalf("batch-input.jsonl 作成失敗: %v", err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)

	modelName := os.Getenv("OPENAI_MODEL")
	if modelName == "" {
		modelName = "gpt-5"
	}

	for _, q := range *questions {
		if *targetQuestionID != 0 && q.ID != int(*targetQuestionID) {
			continue
		}

		log.Printf("問題 ID=%d のヒントを生成中...", q.ID)

		labels := q.Labels
		if len(labels) == 0 {
			log.Printf("問題 ID=%d にはラベルがありません", q.ID)
			continue
		}

		correctAnswers := make(map[string]string)
		for _, am := range q.AnswerMappings {
			correctAnswers[am.Label.LabelCode] = am.Choice.ChoiceCode
		}

		combinations := generateAllCombinations(labels, correctAnswers)

		for i, combo := range combinations {
			if i%10 == 0 {
				log.Printf("  進捗: %d/%d", i, len(combinations))
			}

			// ヒントを生成
			prompt, err := service.BuildPromptForQuestion(&q, combo)
			if err != nil {
				log.Printf("  プロンプト生成エラー: %v", err)
				continue
			}
			stateKey := generateStateKey(combo)
			customID := fmt.Sprintf("q%d_%s", q.ID, stateKey)

			line := BatchLine{
				CustomID: customID,
				Method:   "POST",
				URL:      "/v1/chat/completions",
				Body: map[string]interface{}{
					"model": modelName,
					"messages": []map[string]string{
						{
							"role":    "user",
							"content": prompt,
						},
					},
				},
			}

			if err := enc.Encode(&line); err != nil {
				log.Printf("  JSONL書き込みエラー: %v", err)
				continue
			}
		}

		log.Printf("問題 ID=%d のヒント生成完了", q.ID)
	}

	log.Println("✅ すべてのbatch生成完了")
}

// 各ラベルについて「埋まっている」「埋まっていない」の2通り
// ただし、全て埋まっている状態は除外
func generateAllCombinations(labels []model.Label, correctAnswers map[string]string) []map[string]*string {
	n := len(labels)
	total := 1 << n

	combinations := make([]map[string]*string, 0, total)

	for i := 0; i < total; i++ {
		combo := make(map[string]*string)
		filledCount := 0

		for j, label := range labels {
			if (i>>j)&1 == 1 {
				if choiceCode, ok := correctAnswers[label.LabelCode]; ok {
					combo[label.LabelCode] = &choiceCode
				} else {
					dummy := ""
					combo[label.LabelCode] = &dummy
					log.Printf("警告: ラベル %s の正解が見つかりません", label.LabelCode)
				}
				filledCount++
			}
		}

		if filledCount == n {
			continue
		}

		combinations = append(combinations, combo)
	}

	return combinations
}

func generateStateKey(answers map[string]*string) string {
	var keys []string
	for k := range answers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return strings.Join(keys, ",")
}
