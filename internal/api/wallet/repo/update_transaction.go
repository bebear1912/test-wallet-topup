package repo

import (
	"context"
	"wallet-topup/internal/entities"
)

func (r *WalletRepository) UpdateTransaction(ctx context.Context, tx *entities.Transaction) error {
	return r.db.WithContext(ctx).Save(tx).Error
}
