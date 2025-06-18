package dto

type ConsumerLimitResponse struct {
	ID         string `json:"id"`
	ConsumerID string `json:"consumer_id"`
	TenorMonth int    `json:"tenor_month"`
	MaxLimit   int64  `json:"max_limit"`
	UsedLimit  int64  `json:"used_limit"`
	Available  int64  `json:"available"`
}
