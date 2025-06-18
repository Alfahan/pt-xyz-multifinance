package usecase

import (
	"context"
	"testing"

	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type mockTransactionRepository struct {
	mock.Mock
}

func (m *mockTransactionRepository) CreateTransaction(ctx context.Context, req dto.TransactionRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *mockTransactionRepository) GetTransactionByID(ctx context.Context, id string) (*dto.TransactionResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*dto.TransactionResponse), args.Error(1)
}

type mockConsumerLimitRepository struct {
	mock.Mock
}

func (m *mockConsumerLimitRepository) GetLimitByConsumerIDAndTenor(ctx context.Context, consumerID string, tenorMonth int) (*domain.ConsumerLimit, error) {
	args := m.Called(ctx, consumerID, tenorMonth)
	return args.Get(0).(*domain.ConsumerLimit), args.Error(1)
}

func (m *mockConsumerLimitRepository) UpdateLimit(ctx context.Context, limit domain.ConsumerLimit) error {
	args := m.Called(ctx, limit)
	return args.Error(0)
}

func TestCreateTransaction_ValidLimit(t *testing.T) {
	transactionRepo := new(mockTransactionRepository)
	consumerLimitRepo := new(mockConsumerLimitRepository)
	transactionUC := usecase.NewTransactionUsecase(transactionRepo, consumerLimitRepo)

	// Mock data
	limit := &domain.ConsumerLimit{
		ID:         "1",
		ConsumerID: "123",
		TenorMonth: 3,
		MaxLimit:   500000,
		UsedLimit:  200000,
	}
	req := dto.TransactionRequest{
		ConsumerID:  "123",
		TenorMonth:  3,
		Installment: 100000,
	}

	// Mock expectations
	consumerLimitRepo.On("GetLimitByConsumerIDAndTenor", mock.Anything, req.ConsumerID, req.TenorMonth).Return(limit, nil)
	consumerLimitRepo.On("UpdateLimit", mock.Anything, mock.MatchedBy(func(l domain.ConsumerLimit) bool {
		return l.UsedLimit == 300000
	})).Return(nil)
	transactionRepo.On("CreateTransaction", mock.Anything, req).Return(nil)

	// Execute usecase
	err := transactionUC.CreateTransaction(context.Background(), req)

	// Assert
	assert.NoError(t, err)
	consumerLimitRepo.AssertExpectations(t)
	transactionRepo.AssertExpectations(t)
}

func TestCreateTransaction_ExceedLimit(t *testing.T) {
	transactionRepo := new(mockTransactionRepository)
	consumerLimitRepo := new(mockConsumerLimitRepository)
	transactionUC := usecase.NewTransactionUsecase(transactionRepo, consumerLimitRepo)

	// Mock data
	limit := &domain.ConsumerLimit{
		ID:         "1",
		ConsumerID: "123",
		TenorMonth: 3,
		MaxLimit:   500000,
		UsedLimit:  450000,
	}
	req := dto.TransactionRequest{
		ConsumerID:  "123",
		TenorMonth:  3,
		Installment: 100000,
	}

	// Mock expectations
	consumerLimitRepo.On("GetLimitByConsumerIDAndTenor", mock.Anything, req.ConsumerID, req.TenorMonth).Return(limit, nil)

	// Execute usecase
	err := transactionUC.CreateTransaction(context.Background(), req)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "limit exceeded", err.Error())
	consumerLimitRepo.AssertExpectations(t)
	transactionRepo.AssertNotCalled(t, "CreateTransaction")
}
