// usecase/log_usecase.go
package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
	"gorm.io/datatypes"
)

type LogCommand struct {
	Sub            *string
	StudentNo6     *string
	QuestionID     *int
	EventName      string
	Answers        map[string]*string
	HintOpenStatus map[string]bool
	Hints          map[string]*string
	AnonID         *string
	ClientAt       *time.Time
}

type ILogUsecase interface {
	SendLog(ctx context.Context, in LogCommand) (int, error)
}

type logUsecase struct {
	logRepo  repository.ILogRepository
	userRepo repository.IUserRepository
}

func NewLogUsecase(lr repository.ILogRepository, ur repository.IUserRepository) ILogUsecase {
	return &logUsecase{logRepo: lr, userRepo: ur}
}

func (uc *logUsecase) SendLog(ctx context.Context, cmd LogCommand) (int, error) {

	if cmd.Sub != nil && cmd.StudentNo6 != nil {
		if _, err := uc.userRepo.GetOrCreateBySub(ctx, cmd.Sub, cmd.StudentNo6); err != nil {
			return 0, err
		}
	}

	details := map[string]any{
		"answers":          cmd.Answers,
		"hint_open_status": cmd.HintOpenStatus,
		"hints":            cmd.Hints,
	}
	raw, _ := json.Marshal(details)

	log := model.OperationLog{
		AnonID:          cmd.AnonID,
		UserSub:         cmd.Sub,
		QuestionId:      derefInt(cmd.QuestionID),
		EventName:       cmd.EventName,
		Details:         datatypes.JSON(raw),
		ClientTimestamp: cmd.ClientAt,
	}
	return uc.logRepo.SaveLog(ctx, &log)
}

func derefInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
