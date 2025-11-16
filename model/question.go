package model

import (
	"fmt"
	"strings"
)

type Question struct {
	ID                   int             `json:"id" gorm:"primaryKey"`
	OrderIndex           int             `json:"order_index"`
	ProblemDescription   string          `json:"problem_description"`
	Question             string          `json:"question"`
	AnswerProcess        string          `json:"answer_process"`
	ClassDiagramImage    string          `json:"image"`
	ClassDiagramPlantUML string          `json:"class_diagram_plantuml"`
	Choices              []Choice        `json:"choices" gorm:"foreignKey:QuestionID"`
	Labels               []Label         `json:"labels" gorm:"foreignKey:QuestionID"`
	AnswerMappings       []AnswerMapping `json:"answer_mappings" gorm:"foreignKey:QuestionID"`
	CreatedAt            int64           `gorm:"autoCreateTime"`
}

func (q *Question) Check(ans map[string]*string) (bool, string) {
	correctAnswers := make(map[string]string)
	for _, am := range q.AnswerMappings {
		correctAnswers[am.Label.LabelCode] = am.Choice.ChoiceCode
	}

	var errs []string

	for label, want := range correctAnswers {
		got := ""
		if v, ok := ans[label]; ok && v != nil {
			got = strings.TrimSpace(*v)
			// 回答があればチェック
			if got != "" && got != want {
				errs = append(errs, fmt.Sprintf("%s の正解は %s（あなた: %s）", label, want, got))
			}
		}
	}

	// 不要なラベルはエラーにする
	for label := range ans {
		if _, ok := correctAnswers[label]; !ok {
			errs = append(errs, fmt.Sprintf("不要なラベル %s が含まれています", label))
		}
	}

	if len(errs) > 0 {
		return false, strings.Join(errs, "\n")
	}
	return true, "現在の回答は途中まで正しいです"
}
