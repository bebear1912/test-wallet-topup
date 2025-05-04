package services

import (
	"context"
	"fmt"
	"log"
	"time"
	"wallet-topup/internal/entities"

	"github.com/google/uuid"
)

func (s *WalletService) ConfirmTransaction(ctx context.Context, transactionID uuid.UUID) (*entities.Transaction, error) {
	log.Println("Starting confirm transaction", "transaction_id", transactionID)

	// Get transaction
	tx, err := s.repo.GetTransaction(ctx, transactionID)
	if err != nil {
		log.Println("Failed to get transaction",
			"transaction_id", transactionID,
			"error", err)
		return nil, err
	}

	// Check if transaction is already completed
	if tx.Status == "completed" {
		log.Println("Transaction already completed",
			"transaction_id", transactionID)
		return nil, fmt.Errorf("this transaction status already completed")
	}

	// Check if transaction is expired and update status
	if tx.Status == "expired" {
		log.Println("Transaction already expired",
			"transaction_id", transactionID)
		return nil, fmt.Errorf("this transaction is expired")
	}

	// Mark as expired if past expiry time
	if time.Now().After(tx.ExpiresAt) {
		tx.Status = "expired"
		if err := s.repo.UpdateTransaction(ctx, tx); err != nil {
			log.Println("Failed to update transaction status to expired",
				"transaction_id", transactionID,
				"error", err)
			return nil, err
		}
		log.Println("Transaction marked as expired",
			"transaction_id", transactionID,
			"expires_at", tx.ExpiresAt)
		return nil, fmt.Errorf("transaction expired")
	}

	// Check if transaction is expired
	if time.Now().After(tx.ExpiresAt) {
		log.Println("Transaction expired",
			"transaction_id", transactionID,
			"expires_at", tx.ExpiresAt)
		tx.Status = "failed"
		if err := s.repo.UpdateTransaction(ctx, tx); err != nil {
			log.Println("Failed to update expired transaction",
				"transaction_id", transactionID,
				"error", err)
			return nil, err
		}
		return tx, nil
	}

	// Get user's current balance
	user, err := s.repo.GetUser(ctx, tx.UserID)
	if err != nil {
		log.Println("Failed to get user",
			"user_id", tx.UserID,
			"error", err)
		return nil, err
	}

	// Update user balance
	if err := s.repo.UpdateUserBalance(ctx, tx.UserID, tx.Amount); err != nil {
		log.Println("Failed to update user balance",
			"user_id", tx.UserID,
			"amount", tx.Amount,
			"error", err)
		return nil, err
	}

	// Update transaction status
	tx.Status = "completed"
	tx.UpdatedAt = time.Now()
	tx.Balance = user.Balance + tx.Amount
	if err := s.repo.UpdateTransaction(ctx, tx); err != nil {
		log.Println("Failed to update transaction status",
			"transaction_id", transactionID,
			"error", err)
		return nil, err
	}

	log.Println("Successfully confirmed transaction",
		"transaction_id", transactionID,
		"user_id", tx.UserID,
		"amount", tx.Amount,
		"new_balance", tx.Balance)

	return tx, nil
}
