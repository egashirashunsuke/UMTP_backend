package model

import (
	"time"

	"gorm.io/datatypes"
)

type OperationLog struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	AnonID          *string        `json:"anon_id" `
	UserSub         *string        `gorm:"index"`
	User            *User          `gorm:"foreignKey:UserSub;references:Sub;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	QuestionId      int            `json:"question_id"`
	EventName       string         `json:"event_name"`
	Details         datatypes.JSON `json:"details"`
	ClientTimestamp *time.Time     `json:"timestamp"`
	ServerTimestamp time.Time      `json:"server_timestamp" gorm:"autoCreateTime"`
}
