package repository

import (
	"github.com/egashirashunsuke/UMTP_backend/model"
	"gorm.io/gorm"
)

type ILogRepository interface {
	SaveLog(log *model.OperationLog) error
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) ILogRepository {
	return &logRepository{db: db}
}
func (r *logRepository) SaveLog(log *model.OperationLog) error {
	if err := r.db.Create(log).Error; err != nil {
		return err
	}
	return nil
}
