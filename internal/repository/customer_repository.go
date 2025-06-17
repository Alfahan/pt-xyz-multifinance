package repository

import (
	"context"
	"database/sql"
	"errors"
	"pt-xyz-multifinance/internal/domain"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *domain.Customer) error
	GetByID(ctx context.Context, id string) (*domain.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	query := `
		INSERT INTO consumers (
			id, nik, full_name, legal_name, birth_place, birth_date, salary, photo_ktp, photo_selfie, created_at, updated_at
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
		)
	`
	_, err := r.db.ExecContext(
		ctx, query,
		customer.ID,
		customer.NIK,
		customer.FullName,
		customer.LegalName,
		customer.BirthPlace,
		customer.BirthDate,
		customer.Salary,
		customer.KTPPhotoURL,    // photo_ktp di SQL
		customer.SelfiePhotoURL, // photo_selfie di SQL
		customer.CreatedAt,
		customer.UpdatedAt,
	)
	return err
}

func (r *customerRepository) GetByID(ctx context.Context, id string) (*domain.Customer, error) {
	query := `
		SELECT id, nik, full_name, legal_name, birth_place, birth_date, salary, photo_ktp, photo_selfie, created_at, updated_at
		FROM consumers
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var customer domain.Customer
	err := row.Scan(
		&customer.ID,
		&customer.NIK,
		&customer.FullName,
		&customer.LegalName,
		&customer.BirthPlace,
		&customer.BirthDate,
		&customer.Salary,
		&customer.KTPPhotoURL,    // photo_ktp di SQL
		&customer.SelfiePhotoURL, // photo_selfie di SQL
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}
