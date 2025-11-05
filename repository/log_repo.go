package repository

import (
	"context"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"gorm.io/gorm"
)

type ILogRepository interface {
	SaveLog(ctx context.Context, log *model.OperationLog) (int, error)
}

type logRepository struct{ db *gorm.DB }

func NewLogRepository(db *gorm.DB) ILogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) SaveLog(ctx context.Context, log *model.OperationLog) (int, error) {
	if err := r.db.WithContext(ctx).Create(log).Error; err != nil {
		return 0, err
	}
	return log.ID, nil
}
