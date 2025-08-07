package model

import (
	"time"

	"github.com/google/uuid"
)

type OperationLog struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	DeviceId        uuid.UUID `json:"device_id"`
	EventName       string    `json:"event_name"`
	Timestamp       string    `json:"timestamp"`
	ServerTimestamp time.Time `json:"server_timestamp" gorm:"autoCreateTime"`
}
