package model

type Question struct {
	ID        int      `json:"id" gorm:"primaryKey"`
	Question  string   `json:"question"`
	Answer    string   `json:"answer"`
	Image     string   `json:"image"`
	Choices   []Choice `json:"choices" gorm:"foreignKey:QuestionID"`
	CreatedAt string   `json:"created_at"`
}
