package usecase

import (
	"encoding/json"
	"time"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type LogData struct {
	DeviceId       string             `json:"device_id"`
	QuestionId     int                `json:"question_id"`
	EventName      string             `json:"event_name"`
	Answers        map[string]*string `json:"answers"`
	HintOpenStatus map[string]bool    `json:"hint_open_status"`
	Hints          map[string]*string `json:"hints"`
	Timestamp      string             `json:"timestamp"`
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
	detailsMap := map[string]interface{}{
		"answers":          logdata.Answers,
		"hint_open_status": logdata.HintOpenStatus,
		"hints":            logdata.Hints,
	}

	detailsJSON, err := json.Marshal(detailsMap)
	if err != nil {
		return err
	}

	op := &model.OperationLog{
		DeviceId:        devUUID,
		QuestionId:      logdata.QuestionId,
		EventName:       logdata.EventName,
		Timestamp:       logdata.Timestamp,
		Details:         datatypes.JSON(detailsJSON),
		ServerTimestamp: time.Now(),
	}

	return lu.lr.SaveLog(op)
}
