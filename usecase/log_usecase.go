package usecase

import (
	"time"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
	"github.com/google/uuid"
)

type LogData struct {
	DeviceId  string `json:"device_id"`
	EventName string `json:"event_name"`
	Timestamp string `json:"timestamp"`
}

type ILogUsecase interface {
	SendLog(logdata *LogData) error
}

type LogUsecase struct {
	lr repository.ILogRepository
}

func NewLogUsecase(lr repository.ILogRepository) ILogUsecase {
	return &LogUsecase{lr: lr}
}

func (lu *LogUsecase) SendLog(logdata *LogData) error {
	devUUID, err := uuid.Parse(logdata.DeviceId)
	if err != nil {
		return err
	}

	op := &model.OperationLog{
		DeviceId:        devUUID,
		EventName:       logdata.EventName,
		Timestamp:       logdata.Timestamp,
		ServerTimestamp: time.Now(),
	}

	return lu.lr.SaveLog(op)
}
