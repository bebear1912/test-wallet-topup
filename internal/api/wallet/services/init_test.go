package services

import (
	"context"
	"testing"
	"time"
	"wallet-topup/internal/api/wallet"
	"wallet-topup/internal/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUser(ctx context.Context, userID uint) (*entities.User, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockRepository) CreateTransaction(ctx context.Context, tx *entities.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

func (m *MockRepository) GetTransaction(ctx context.Context, transactionID uuid.UUID) (*entities.Transaction, error) {
	args := m.Called(ctx, transactionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockRepository) UpdateTransaction(ctx context.Context, tx *entities.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

func (m *MockRepository) UpdateUserBalance(ctx context.Context, userID uint, amount float64) error {
	args := m.Called(ctx, userID, amount)
	return args.Error(0)
}

func TestVerifyTransaction(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	ctx := context.Background()
	req := wallet.VerifyRequest{
		UserID:        1,
		Amount:        100.50,
		PaymentMethod: "credit_card",
	}

	// Mock GetUser to return a user
	mockRepo.On("GetUser", ctx, req.UserID).Return(&entities.User{
		ID:      req.UserID,
		Balance: 1000.00,
	}, nil)

	// Mock CreateTransaction
	mockRepo.On("CreateTransaction", ctx, mock.Anything).Return(nil)

	tx, err := service.VerifyTransaction(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, tx)
	assert.Equal(t, req.UserID, tx.UserID)
	assert.Equal(t, req.Amount, tx.Amount)
	assert.Equal(t, req.PaymentMethod, tx.PaymentMethod)
	assert.Equal(t, "verified", tx.Status)
	assert.True(t, time.Now().Before(tx.ExpiresAt))

	mockRepo.AssertExpectations(t)
}

func TestConfirmTransaction(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	ctx := context.Background()
	transactionID := uuid.New()
	userID := uint(1)
	amount := 100.50

	// Mock GetUser to return a user
	mockRepo.On("GetUser", ctx, userID).Return(&entities.User{
		ID:      userID,
		Balance: 1000.00,
	}, nil)

	// Mock GetTransaction to return a valid transaction
	mockRepo.On("GetTransaction", ctx, transactionID).Return(&entities.Transaction{
		ID:        transactionID,
		UserID:    userID,
		Amount:    amount,
		Status:    "verified",
		ExpiresAt: time.Now().Add(1 * time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	// Mock UpdateUserBalance
	mockRepo.On("UpdateUserBalance", ctx, userID, amount).Return(nil)

	// Mock UpdateTransaction
	mockRepo.On("UpdateTransaction", ctx, mock.Anything).Return(nil)

	tx, err := service.ConfirmTransaction(ctx, transactionID)
	assert.NoError(t, err)
	assert.NotNil(t, tx)
	assert.Equal(t, "completed", tx.Status)

	mockRepo.AssertExpectations(t)
}
