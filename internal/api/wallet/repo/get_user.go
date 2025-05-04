package repo

import (
	"context"
	"wallet-topup/internal/entities"
)

func (r *WalletRepository) GetUser(ctx context.Context, userID uint) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
