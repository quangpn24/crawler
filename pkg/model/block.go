package model

import (
	"github.com/google/uuid"
)

type Block struct {
	BaseModel
	Data    string    `json:"data"`
	HTMLTag string    `json:"html_tag"`
	PostID  uuid.UUID `json:"post_id" gorm:"type:uuid"`
	Rank    int       `json:"rank"`
}
