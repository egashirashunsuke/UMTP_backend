package dto

type LogRequest struct {
	QuestionID     *int               `json:"question_id"`
	StudentID      *string            `json:"student_id"`
	EventName      string             `json:"event_name"`
	Answers        map[string]*string `json:"answers"`
	HintOpenStatus map[string]bool    `json:"hint_open_status"`
	Hints          map[string]*string `json:"hints"`
	AnonID         *string            `json:"anon_id"`
	Timestamp      *string            `json:"timestamp"` // ISO8601 (RFC3339) 文字列
}
