package dto

type TransactionRequest struct {
	ConsumerID     string `json:"consumer_id" validate:"required,uuid"`
	TenorMonth     int    `json:"tenor_month" validate:"required,oneof=1 2 3 6"`
	OTR            int64  `json:"otr" validate:"required,min=1"`
	AdminFee       int64  `json:"admin_fee" validate:"required,min=0"`
	Installment    int64  `json:"installment" validate:"required,min=1"`
	Interest       int64  `json:"interest" validate:"required,min=0"`
	AssetName      string `json:"asset_name" validate:"required"`
	ContractNumber string `json:"contract_number" validate:"required"`
}

type TransactionResponse struct {
	ID             string `json:"id"`
	ContractNumber string `json:"contract_number"`
	ConsumerID     string `json:"consumer_id"`
	TenorMonth     int    `json:"tenor_month"`
	OTR            int64  `json:"otr"`
	AdminFee       int64  `json:"admin_fee"`
	Installment    int64  `json:"installment"`
	Interest       int64  `json:"interest"`
	AssetName      string `json:"asset_name"`
	CreatedAt      string `json:"created_at"`
}
