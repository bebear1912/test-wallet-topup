package repo

import (
	"wallet-topup/internal/api/wallet"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) wallet.Repository {
	return &WalletRepository{db: db}
}
