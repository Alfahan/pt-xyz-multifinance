package usecase

import (
	"context"
	"errors"
	"time"

	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/repository"

	"github.com/google/uuid"
)

type CustomerUsecase interface {
	Create(ctx context.Context, req *dto.CreateCustomerRequest) (*dto.CustomerResponse, error)
	GetByID(ctx context.Context, id string) (*dto.CustomerResponse, error)
}

type customerUsecase struct {
	repo repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo: repo}
}

func (u *customerUsecase) Create(ctx context.Context, req *dto.CreateCustomerRequest) (*dto.CustomerResponse, error) {
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, errors.New("invalid birth date format, use YYYY-MM-DD")
	}
	now := time.Now()
	customer := &domain.Customer{
		ID:             uuid.New().String(),
		NIK:            req.NIK,
		FullName:       req.FullName,
		LegalName:      req.LegalName,
		BirthPlace:     req.BirthPlace,
		BirthDate:      birthDate,
		Salary:         req.Salary,
		KTPPhotoURL:    req.KTPPhotoURL,
		SelfiePhotoURL: req.SelfiePhotoURL,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	err = u.repo.Create(ctx, customer)
	if err != nil {
		return nil, err
	}
	return customerToResponse(customer), nil
}

func (u *customerUsecase) GetByID(ctx context.Context, id string) (*dto.CustomerResponse, error) {
	customer, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}
	return customerToResponse(customer), nil
}

func customerToResponse(customer *domain.Customer) *dto.CustomerResponse {
	return &dto.CustomerResponse{
		ID:             customer.ID,
		NIK:            customer.NIK,
		FullName:       customer.FullName,
		LegalName:      customer.LegalName,
		BirthPlace:     customer.BirthPlace,
		BirthDate:      customer.BirthDate.Format("2006-01-02"),
		Salary:         customer.Salary,
		KTPPhotoURL:    customer.KTPPhotoURL,
		SelfiePhotoURL: customer.SelfiePhotoURL,
		CreatedAt:      customer.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      customer.UpdatedAt.Format(time.RFC3339),
	}
}
