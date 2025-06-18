package usecase

import (
	"context"
	"errors"
	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/repository"
	"sync"
)

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, req dto.TransactionRequest) error
	GetTransactionByID(ctx context.Context, id string) (*dto.TransactionResponse, error)
}

type transactionUsecase struct {
	transactionRepo   repository.TransactionRepository
	consumerLimitRepo repository.ConsumerLimitRepository
	mu                sync.Mutex // Mutex untuk concurrency handling
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository, consumerLimitRepo repository.ConsumerLimitRepository) TransactionUsecase {
	return &transactionUsecase{transactionRepo: transactionRepo, consumerLimitRepo: consumerLimitRepo}
}

func (uc *transactionUsecase) CreateTransaction(ctx context.Context, req dto.TransactionRequest) error {
	uc.mu.Lock() // Lock untuk mencegah race condition
	defer uc.mu.Unlock()

	// Check consumer limit
	limit, err := uc.consumerLimitRepo.GetLimitByConsumerIDAndTenor(ctx, req.ConsumerID, req.TenorMonth)
	if err != nil {
		return errors.New("failed to fetch consumer limit")
	}

	if limit.UsedLimit+req.Installment > limit.MaxLimit {
		return errors.New("limit exceeded")
	}

	// Update consumer limit
	limit.UsedLimit += req.Installment
	err = uc.consumerLimitRepo.UpdateLimit(ctx, *limit) // Mendereferensikan pointer
	if err != nil {
		return errors.New("failed to update consumer limit")
	}

	// Create transaction
	err = uc.transactionRepo.CreateTransaction(ctx, req)
	if err != nil {
		return errors.New("failed to create transaction")
	}

	return nil
}

func (uc *transactionUsecase) GetTransactionByID(ctx context.Context, id string) (*dto.TransactionResponse, error) {
	return uc.transactionRepo.GetTransactionByID(ctx, id)
}
