package key

import (
	"crypto/rand"
	"encoding/hex"
	"gorm.io/gorm"
	"time"
)

type Key struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
}

type Struct struct {
	*gorm.DB
}

func NewDb(db *gorm.DB) *Struct {
	value := Struct{
		db,
	}
	return &value
}

func generateKey() (*string, error) {
	bytes := make([]byte, 32) // 32 bytes will give us a 64 character hex string
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	key := hex.EncodeToString(bytes)
	return &key, nil
}
