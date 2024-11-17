package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Key struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Key       string `gorm:"not null, unique"`
	CreatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (k *Key) TableName() string {
	return "keys"
}
