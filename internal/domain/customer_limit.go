package domain

import "time"

// ConsumerLimit adalah model domain untuk entitas batas limit konsumen
type ConsumerLimit struct {
	ID         string    `json:"id"`
	ConsumerID string    `json:"consumer_id"`
	TenorMonth int       `json:"tenor_month"`
	MaxLimit   int64     `json:"max_limit"`
	UsedLimit  int64     `json:"used_limit"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
