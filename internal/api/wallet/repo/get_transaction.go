package repo

import (
	"context"
	"wallet-topup/internal/entities"

	"github.com/google/uuid"
)

func (r *WalletRepository) GetTransaction(ctx context.Context, transactionID uuid.UUID) (*entities.Transaction, error) {
	var tx entities.Transaction
	if err := r.db.WithContext(ctx).First(&tx, "id = ?", transactionID).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}
