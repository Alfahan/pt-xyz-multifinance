package domain

import "time"

// Transaction adalah model domain untuk entitas transaksi
type Transaction struct {
	ID             string    `json:"id"`
	ContractNumber string    `json:"contract_number"`
	ConsumerID     string    `json:"consumer_id"`
	TenorMonth     int       `json:"tenor_month"`
	OTR            int64     `json:"otr"`
	AdminFee       int64     `json:"admin_fee"`
	Installment    int64     `json:"installment"`
	Interest       int64     `json:"interest"`
	AssetName      string    `json:"asset_name"`
	CreatedAt      time.Time `json:"created_at"`
}
