package schemas

import "time"

type Pack struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Pack) TableName() string {
	return "packs"
}
