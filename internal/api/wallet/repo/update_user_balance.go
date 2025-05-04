package repo

import (
	"context"
	"wallet-topup/internal/entities"

	"gorm.io/gorm"
)

func (r *WalletRepository) UpdateUserBalance(ctx context.Context, userID uint, amount float64) error {
	return r.db.WithContext(ctx).Model(&entities.User{}).
		Where("id = ?", userID).
		Update("balance", gorm.Expr("balance + ?", amount)).
		Error
}
