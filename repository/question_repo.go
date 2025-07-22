package repository

import (
	"fmt"
	"log"

	"github.com/egashirashunsuke/UMTP_backend/dto"
	"github.com/egashirashunsuke/UMTP_backend/model"
	"gorm.io/gorm"
)

type IQuestionRepository interface {
	GetQuestionByID(id int) (*model.Question, error)
	GetAllQuestions() (*[]model.Question, error)
	CreateQuestionWithAssociations(in *dto.CreateQuestionDTO) error
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) GetQuestionByID(id int) (*model.Question, error) {
	var question model.Question
	if err := r.db.Preload("Choices").Preload("AnswerMappings.Choice").
		Preload("AnswerMappings.Label").
		Preload("Choices").
		Preload("Labels").First(&question, id).Error; err != nil {
		return nil, err
	}
	log.Printf("[repo] loaded: choices=%d labels=%d am=%d",
		len(question.Choices), len(question.Labels), len(question.AnswerMappings))

	for i, am := range question.AnswerMappings {
		log.Printf("[repo] AM[%d] label=%s choice=%s",
			i, am.Label.LabelCode, am.Choice.ChoiceCode)
	}

	return &question, nil
}

func (r *questionRepository) GetAllQuestions() (*[]model.Question, error) {
	var questions []model.Question
	if err := r.db.Preload("Choices").Find(&questions).Error; err != nil {
		return nil, fmt.Errorf("failed to get all questions: %w", err)
	}
	if len(questions) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &questions, nil
}

func (r *questionRepository) CreateQuestionWithAssociations(in *dto.CreateQuestionDTO) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		q := model.Question{
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
			ch := model.Choice{
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
			lb := model.Label{
				QuestionID: q.ID,
				LabelCode:  li.LabelCode,
			}
			if err := tx.Create(&lb).Error; err != nil {
				return err
			}
			labelMap[li.LabelCode] = lb.ID
		}
		log.Printf("Created question %d with choices %v and labels %v", q.ID, choiceMap, labelMap)

		var mappings []model.AnswerMapping
		for _, am := range in.AnswerMappings {
			cid, okc := choiceMap[am.ChoiceCode]
			lid, okl := labelMap[am.LabelCode]
			if !okc || !okl {
				return fmt.Errorf("invalid mapping: %s→%s", am.ChoiceCode, am.LabelCode)
			}
			mappings = append(mappings, model.AnswerMapping{
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
