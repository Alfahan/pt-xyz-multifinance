package repository

import (
	"context"
	"database/sql"
	"pt-xyz-multifinance/internal/handler/dto"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, req dto.TransactionRequest) error
	GetTransactionByID(ctx context.Context, id string) (*dto.TransactionResponse, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (repo *transactionRepository) CreateTransaction(ctx context.Context, req dto.TransactionRequest) error {
	query := `INSERT INTO transactions (id, contract_number, consumer_id, tenor_month, otr, admin_fee, installment, interest, asset_name, created_at)
            VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP)`
	_, err := repo.db.ExecContext(ctx, query, req.ContractNumber, req.ConsumerID, req.TenorMonth, req.OTR, req.AdminFee, req.Installment, req.Interest, req.AssetName)
	return err
}

func (repo *transactionRepository) GetTransactionByID(ctx context.Context, id string) (*dto.TransactionResponse, error) {
	query := `SELECT id, contract_number, consumer_id, tenor_month, otr, admin_fee, installment, interest, asset_name, created_at FROM transactions WHERE id = $1`
	row := repo.db.QueryRowContext(ctx, query, id)

	var transaction dto.TransactionResponse
	err := row.Scan(&transaction.ID, &transaction.ContractNumber, &transaction.ConsumerID, &transaction.TenorMonth, &transaction.OTR, &transaction.AdminFee, &transaction.Installment, &transaction.Interest, &transaction.AssetName, &transaction.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
