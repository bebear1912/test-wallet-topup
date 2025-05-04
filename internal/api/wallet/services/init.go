package services

import (
	"wallet-topup/internal/api/wallet"
)

type WalletService struct {
	repo wallet.Repository
}

func NewService(repo wallet.Repository) wallet.Service {
	return &WalletService{repo: repo}
}
