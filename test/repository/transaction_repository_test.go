package repository

import (
	"context"
	"testing"

	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	db, mock, _ := sqlmock.New()
	repo := repository.NewTransactionRepository(db)

	req := dto.TransactionRequest{
		ContractNumber: "123",
		ConsumerID:     "456",
		TenorMonth:     3,
		OTR:            1000000,
		AdminFee:       10000,
		Installment:    300000,
		Interest:       50000,
		AssetName:      "Motor",
	}

	mock.ExpectExec(`INSERT INTO transactions`).WithArgs(
		req.ContractNumber, req.ConsumerID, req.TenorMonth, req.OTR, req.AdminFee, req.Installment, req.Interest, req.AssetName,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateTransaction(context.Background(), req)
	assert.NoError(t, err)
	mock.ExpectationsWereMet()
}
