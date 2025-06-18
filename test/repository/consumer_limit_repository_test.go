package repository

import (
	"context"
	"pt-xyz-multifinance/internal/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetLimitByConsumerIDAndTenor(t *testing.T) {
	db, mock, _ := sqlmock.New()
	repo := repository.NewConsumerLimitRepository(db)

	consumerID := "123"
	tenorMonth := 3

	createdAt, _ := time.Parse("2006-01-02", "2025-06-01")
	updatedAt, _ := time.Parse("2006-01-02", "2025-06-01")

	rows := sqlmock.NewRows([]string{"id", "consumer_id", "tenor_month", "max_limit", "used_limit", "created_at", "updated_at"}).
		AddRow("1", consumerID, tenorMonth, 500000, 200000, createdAt, updatedAt)

	mock.ExpectQuery(`SELECT id, consumer_id, tenor_month, max_limit, used_limit, created_at, updated_at FROM consumer_limits`).
		WithArgs(consumerID, tenorMonth).WillReturnRows(rows)

	limit, err := repo.GetLimitByConsumerIDAndTenor(context.Background(), consumerID, tenorMonth)
	assert.NoError(t, err)
	assert.Equal(t, consumerID, limit.ConsumerID)
	assert.Equal(t, tenorMonth, limit.TenorMonth)
	mock.ExpectationsWereMet()
}
