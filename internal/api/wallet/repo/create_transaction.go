package repo

import (
	"context"
	"wallet-topup/internal/entities"
)

func (r *WalletRepository) CreateTransaction(ctx context.Context, tx *entities.Transaction) error {
	return r.db.WithContext(ctx).Create(tx).Error
}
