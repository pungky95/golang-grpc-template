package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string
	DisplayName string
	CreatedAt   time.Time      `json:"created_at" gorm:"not null;type:timestamp;autoCreateTime;column:created_at"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null;type:timestamp;autoUpdateTime;column:updated_at"`
	DeletedAt   gorm.DeletedAt ` gorm:"index"`
}

func (*Role) TableName() string {
	return "roles"
}
