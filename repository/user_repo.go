package repository

import (
	"context"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	FindBySub(ctx context.Context, sub string) (*model.User, error)
	GetOrCreateBySub(ctx context.Context, sub string, studentNo *string) (*model.User, error)
}

type userRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindBySub(ctx context.Context, sub string) (*model.User, error) {
	var u model.User
	if err := r.db.WithContext(ctx).
		First(&u, "sub = ?", sub).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetOrCreateBySub(ctx context.Context, sub string, studentNo *string) (*model.User, error) {
	u := model.User{Sub: sub, StudentNo: studentNo}

	// INSERT ... ON CONFLICT (sub) DO UPDATE SET student_no = EXCLUDED.student_no
	//   WHERE users.student_no IS NULL
	onConflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "sub"}},
	}

	if studentNo != nil {
		onConflict.DoUpdates = clause.AssignmentColumns([]string{"student_no"})
		onConflict.Where = clause.Where{
			Exprs: []clause.Expression{
				gorm.Expr("users.student_no IS NULL"),
			},
		}
	} else {
		onConflict.DoNothing = true
	}

	if err := r.db.WithContext(ctx).
		Clauses(onConflict).
		Create(&u).Error; err != nil {
		return nil, err
	}

	return r.FindBySub(ctx, sub)
}
