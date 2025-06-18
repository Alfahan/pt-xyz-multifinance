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

type ConsumerUsecase interface {
	Create(ctx context.Context, req *dto.CreateConsumerRequest) (*dto.ConsumerResponse, error)
	GetByID(ctx context.Context, id string) (*dto.ConsumerResponse, error)
}

type consumerUsecase struct {
	repo repository.ConsumerRepository
}

func NewConsumerUsecase(repo repository.ConsumerRepository) ConsumerUsecase {
	return &consumerUsecase{repo: repo}
}

func (u *consumerUsecase) Create(ctx context.Context, req *dto.CreateConsumerRequest) (*dto.ConsumerResponse, error) {
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, errors.New("invalid birth date format, use YYYY-MM-DD")
	}
	now := time.Now()
	consumer := &domain.Consumer{
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
	err = u.repo.Create(ctx, consumer)
	if err != nil {
		return nil, err
	}
	return consumerToResponse(consumer), nil
}

func (u *consumerUsecase) GetByID(ctx context.Context, id string) (*dto.ConsumerResponse, error) {
	consumer, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if consumer == nil {
		return nil, errors.New("consumer not found")
	}
	return consumerToResponse(consumer), nil
}

func consumerToResponse(consumer *domain.Consumer) *dto.ConsumerResponse {
	return &dto.ConsumerResponse{
		ID:             consumer.ID,
		NIK:            consumer.NIK,
		FullName:       consumer.FullName,
		LegalName:      consumer.LegalName,
		BirthPlace:     consumer.BirthPlace,
		BirthDate:      consumer.BirthDate.Format("2006-01-02"),
		Salary:         consumer.Salary,
		KTPPhotoURL:    consumer.KTPPhotoURL,
		SelfiePhotoURL: consumer.SelfiePhotoURL,
		CreatedAt:      consumer.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      consumer.UpdatedAt.Format(time.RFC3339),
	}
}
