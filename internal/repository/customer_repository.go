package repository

import (
	"context"
	"database/sql"
	"errors"
	"pt-xyz-multifinance/internal/domain"
)

type ConsumerRepository interface {
	Create(ctx context.Context, consumer *domain.Consumer) error
	GetByID(ctx context.Context, id string) (*domain.Consumer, error)
}

type consumerRepository struct {
	db *sql.DB
}

func NewConsumerRepository(db *sql.DB) ConsumerRepository {
	return &consumerRepository{db: db}
}

func (r *consumerRepository) Create(ctx context.Context, consumer *domain.Consumer) error {
	query := `
		INSERT INTO consumers (
			id, nik, full_name, legal_name, birth_place, birth_date, salary, photo_ktp, photo_selfie, created_at, updated_at
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
		)
	`
	_, err := r.db.ExecContext(
		ctx, query,
		consumer.ID,
		consumer.NIK,
		consumer.FullName,
		consumer.LegalName,
		consumer.BirthPlace,
		consumer.BirthDate,
		consumer.Salary,
		consumer.KTPPhotoURL,    // photo_ktp di SQL
		consumer.SelfiePhotoURL, // photo_selfie di SQL
		consumer.CreatedAt,
		consumer.UpdatedAt,
	)
	return err
}

func (r *consumerRepository) GetByID(ctx context.Context, id string) (*domain.Consumer, error) {
	query := `
		SELECT id, nik, full_name, legal_name, birth_place, birth_date, salary, photo_ktp, photo_selfie, created_at, updated_at
		FROM consumers
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var consumer domain.Consumer
	err := row.Scan(
		&consumer.ID,
		&consumer.NIK,
		&consumer.FullName,
		&consumer.LegalName,
		&consumer.BirthPlace,
		&consumer.BirthDate,
		&consumer.Salary,
		&consumer.KTPPhotoURL,    // photo_ktp di SQL
		&consumer.SelfiePhotoURL, // photo_selfie di SQL
		&consumer.CreatedAt,
		&consumer.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &consumer, nil
}
