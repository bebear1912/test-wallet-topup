package wallet

import (
	"context"
	"wallet-topup/internal/entities"

	"github.com/google/uuid"
)

type Service interface {
	VerifyTransaction(ctx context.Context, req VerifyRequest) (*entities.Transaction, error)
	ConfirmTransaction(ctx context.Context, transactionID uuid.UUID) (*entities.Transaction, error)
}

type Repository interface {
	GetUser(ctx context.Context, userID uint) (*entities.User, error)
	CreateTransaction(ctx context.Context, tx *entities.Transaction) error
	GetTransaction(ctx context.Context, transactionID uuid.UUID) (*entities.Transaction, error)
	UpdateTransaction(ctx context.Context, tx *entities.Transaction) error
	UpdateUserBalance(ctx context.Context, userID uint, amount float64) error
}

type VerifyRequest struct {
	UserID        uint    `json:"user_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
}
