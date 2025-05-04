package services

import (
	"context"
	"log"
	"time"
	"wallet-topup/internal/api/wallet"
	"wallet-topup/internal/entities"

	"github.com/google/uuid"
)

func (s *WalletService) VerifyTransaction(ctx context.Context, req wallet.VerifyRequest) (*entities.Transaction, error) {
	log.Println("Starting verify transaction",
		"user_id", req.UserID,
		"amount", req.Amount,
		"payment_method", req.PaymentMethod)

	// Check if user exists
	if _, err := s.repo.GetUser(ctx, req.UserID); err != nil {
		log.Println("Failed to get user",
			"user_id", req.UserID,
			"error", err)
		return nil, err
	}

	// Create transaction
	tx := &entities.Transaction{
		ID:            uuid.New(),
		UserID:        req.UserID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		Status:        "verified",
		ExpiresAt:     time.Now().Add(15 * time.Minute),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.repo.CreateTransaction(ctx, tx); err != nil {
		log.Println("Failed to create transaction",
			"transaction_id", tx.ID,
			"error", err)
		return nil, err
	}

	log.Println("Successfully verified transaction",
		"transaction_id", tx.ID,
		"user_id", tx.UserID,
		"amount", tx.Amount)

	return tx, nil
}
