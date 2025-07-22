// api/dto/question.go
package dto

type CreateQuestionDTO struct {
	ProblemDescription   string                   `json:"problem_description"`
	Question             string                   `json:"question"`
	AnswerProcess        string                   `json:"answer_process"`
	ClassDiagramImage    string                   `json:"image"`
	ClassDiagramPlantUML string                   `json:"class_diagram_plantuml"`
	Choices              []CreateChoiceDTO        `json:"choices" gorm:"foreignKey:QuestionID"`
	Labels               []CreateLabelDTO         `json:"labels" gorm:"foreignKey:QuestionID"`
	AnswerMappings       []CreateAnswerMappingDTO `json:"answer_mappings" gorm:"foreignKey:QuestionID"`
}

type CreateChoiceDTO struct {
	ChoiceCode string `json:"choice_code"`
	ChoiceText string `json:"choice_text"`
}

type CreateLabelDTO struct {
	LabelCode string `json:"label_code"`
}

type CreateAnswerMappingDTO struct {
	ChoiceCode string `json:"choice_code"`
	LabelCode  string `json:"label_code"`
}
