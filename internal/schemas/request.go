package schemas

import (
	"encoding/json"
	"time"
)

type RequestApi struct {
	ID         uint            `gorm:"primaryKey"`
	UserID     uint            `gorm:"not null"`
	ClientIP   string          `gorm:"not null"`
	Key        string          `gorm:"not null"`
	Endpoint   string          `gorm:"not null"`
	Method     string          `gorm:"not null"`
	Params     json.RawMessage `gorm:"null"`
	Headers    json.RawMessage `gorm:"null"`
	StatusCode int             `gorm:"not null"`
	Body       json.RawMessage `gorm:"null"`
	Content    json.RawMessage `gorm:"null"`
	Error      string          `gorm:"null"`
	CreatedAt  time.Time
}

func (r *RequestApi) TableName() string {
	return "requests"
}
