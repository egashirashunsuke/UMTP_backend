package model

import (
	"fmt"
	"strings"
)

type Question struct {
	ID                   int             `json:"id" gorm:"primaryKey"`
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
	// 正解マップ作成
	correct := make(map[string]string) // label_code -> choice_code
	for _, am := range q.AnswerMappings {
		correct[am.Label.LabelCode] = am.Choice.ChoiceCode
	}

	var errs []string

	// 必須ラベルが全部答えられているか & 内容一致確認
	for label, want := range correct {
		got := ""
		if v, ok := ans[label]; ok && v != nil {
			got = strings.TrimSpace(*v)
		} else {
			errs = append(errs, fmt.Sprintf("%s が未回答です", label))
			continue
		}
		if got != want {
			errs = append(errs, fmt.Sprintf("%s の正解は %s（あなた: %s）", label, want, got))
		}
	}

	// 余分なラベルが来ていないか
	for label := range ans {
		if _, ok := correct[label]; !ok {
			errs = append(errs, fmt.Sprintf("不要なラベル %s が含まれています", label))
		}
	}

	if len(errs) > 0 {
		return false, strings.Join(errs, "\n")
	}
	return true, "正解！"
}
