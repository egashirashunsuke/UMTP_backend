package model

import (
	"fmt"
	"log"

	"github.com/egashirashunsuke/UMTP_backend/dto"
	"gorm.io/gorm"
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

func GetQuestionByID(db *gorm.DB, id int) (*Question, error) {
	var question Question
	if err := db.Preload("Choices").First(&question, id).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func CreateQuestionWithAssociations(db *gorm.DB, in *dto.CreateQuestionDTO) error {
	return db.Transaction(func(tx *gorm.DB) error {
		q := Question{
			ProblemDescription:   in.ProblemDescription,
			Question:             in.Question,
			ClassDiagramImage:    in.ClassDiagramImage,
			ClassDiagramPlantUML: in.ClassDiagramPlantUML,
		}
		if err := tx.Create(&q).Error; err != nil {
			return err
		}

		// 2) Choices を Insert & code→ID マップ作成
		choiceMap := make(map[string]int)
		for _, ci := range in.Choices {
			ch := Choice{
				QuestionID: q.ID,
				ChoiceCode: ci.ChoiceCode,
				ChoiceText: ci.ChoiceText,
			}
			if err := tx.Create(&ch).Error; err != nil {
				return err
			}
			choiceMap[ci.ChoiceCode] = ch.ID
		}

		labelMap := make(map[string]int)
		for _, li := range in.Labels {
			lb := Label{
				QuestionID: q.ID,
				LabelCode:  li.LabelCode,
			}
			if err := tx.Create(&lb).Error; err != nil {
				return err
			}
			labelMap[li.LabelCode] = lb.ID
		}
		log.Printf("Created question %d with choices %v and labels %v", q.ID, choiceMap, labelMap)

		var mappings []AnswerMapping
		for _, am := range in.AnswerMappings {
			cid, okc := choiceMap[am.ChoiceCode]
			lid, okl := labelMap[am.LabelCode]
			if !okc || !okl {
				return fmt.Errorf("invalid mapping: %s→%s", am.ChoiceCode, am.LabelCode)
			}
			mappings = append(mappings, AnswerMapping{
				QuestionID: q.ID,
				ChoiceID:   cid,
				LabelID:    lid,
			})
		}
		if len(mappings) > 0 {
			if err := tx.Create(&mappings).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (q *Question) Check(ans map[string]*string) (bool, string) {
	// 正解マッピングを作成
	correct := make(map[string]string) // label_code -> choice_code
	for _, am := range q.AnswerMappings {
		// ChoiceとLabelをPreloadしておく必要あり
		correct[am.Label.LabelCode] = am.Choice.ChoiceCode
	}

	// 回答を比較
	for label, userChoicePtr := range ans {
		userChoice := ""
		if userChoicePtr != nil {
			userChoice = *userChoicePtr
		}
		if correctChoice, ok := correct[label]; !ok || userChoice != correctChoice {
			return false, "不正解" // 不正解
		}
	}

	return true, "正解" // 全て一致なら正解
}
