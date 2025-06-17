package repository

import (
	"context"
	"database/sql"
	"pt-xyz-multifinance/internal/domain"
)

type ConsumerLimitRepository interface {
	GetLimitByConsumerIDAndTenor(ctx context.Context, consumerID string, tenorMonth int) (*domain.ConsumerLimit, error)
	UpdateLimit(ctx context.Context, limit domain.ConsumerLimit) error
}

type consumerLimitRepository struct {
	db *sql.DB
}

func NewConsumerLimitRepository(db *sql.DB) ConsumerLimitRepository {
	return &consumerLimitRepository{db: db}
}

func (repo *consumerLimitRepository) GetLimitByConsumerIDAndTenor(ctx context.Context, consumerID string, tenorMonth int) (*domain.ConsumerLimit, error) {
	query := `SELECT id, consumer_id, tenor_month, max_limit, used_limit, created_at, updated_at FROM consumer_limits WHERE consumer_id = $1 AND tenor_month = $2`
	row := repo.db.QueryRowContext(ctx, query, consumerID, tenorMonth)

	var limit domain.ConsumerLimit
	err := row.Scan(&limit.ID, &limit.ConsumerID, &limit.TenorMonth, &limit.MaxLimit, &limit.UsedLimit, &limit.CreatedAt, &limit.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &limit, nil
}

func (repo *consumerLimitRepository) UpdateLimit(ctx context.Context, limit domain.ConsumerLimit) error {
	query := `UPDATE consumer_limits SET used_limit = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`
	_, err := repo.db.ExecContext(ctx, query, limit.UsedLimit, limit.ID)
	return err
}
