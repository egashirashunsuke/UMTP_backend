package dto

type LogRequest struct {
	QuestionID     *int               `json:"question_id"`      // number | undefined
	StudentID      *string            `json:"student_id"`       // string | false | undefined
	EventName      string             `json:"event_name"`       // 必須
	Answers        map[string]*string `json:"answers"`          // 任意
	HintOpenStatus map[string]bool    `json:"hint_open_status"` // keys: "1","2",...
	Hints          map[string]*string `json:"hints"`            // keys: "1","2",...
	HintIndex      *int               `json:"hintIndex"`        // 任意
	Useful         *int               `json:"useful"`           // 任意
	Comment        *string            `json:"comment"`          // 任意
	AnonID         *string            `json:"anon_id"`          // 任意（匿名トラッキング用）
	Timestamp      *string            `json:"timestamp"`        // ISO8601 (RFC3339) 文字列
}
