package model

import "time"

type User struct {
	Sub       *string `json:"id" gorm:"primaryKey"`
	StudentNo *string
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Logs []OperationLog `gorm:"foreignKey:UserSub;references:Sub"`
}
