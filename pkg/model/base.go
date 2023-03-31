package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID       `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatorID *uuid.UUID      `json:"creator_id,omitempty" gorm:"type:uuid;"`
	UpdaterID *uuid.UUID      `json:"updater_id,omitempty" gorm:"type:uuid;"`
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" swaggertype:"string"`
}
