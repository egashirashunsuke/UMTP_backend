package main

import (
	"context"
	"encoding/json"
	"log"
	"sort"
	"strings"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
	"github.com/egashirashunsuke/UMTP_backend/service"
)

func main() {
	db := model.DBConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// すべての問題を取得
	qRepo := repository.NewQuestionRepository(db)
	questions, err := qRepo.GetAllQuestions()
	if err != nil {
		log.Fatalf("問題取得失敗: %v", err)
	}
	log.Printf("問題数: %d", len(*questions))

	hintGen := service.NewHintsService()

	ctx := context.Background()

	for _, q := range *questions {
		if q.ID != 4 {
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
			hints, err := hintGen.Generate(ctx, &q, combo)
			if err != nil {
				log.Printf("  エラー: %v", err)
				continue
			}

			hintsJSON, err := json.Marshal(hints)
			if err != nil {
				log.Printf("  JSON変換エラー: %v", err)
				continue
			}

			stateKey := generateStateKey(combo)

			hint := model.Hint{
				QuestionID:   q.ID,
				AnswersState: stateKey,
				Hints:        string(hintsJSON),
			}

			if err := db.Where("question_id = ? AND answers_state = ?", q.ID, stateKey).
				FirstOrCreate(&hint).Error; err != nil {
				log.Printf("  保存エラー: %v", err)
			}
		}

		log.Printf("問題 ID=%d のヒント生成完了", q.ID)
	}

	log.Println("✅ すべてのヒント生成完了")
}

// 各ラベルについて「埋まっている」「埋まっていない」の2通り
// ただし、全て埋まっている状態と、最後の一つだけ埋まっていない状態は除外
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

		if filledCount == n || filledCount == n-1 {
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
