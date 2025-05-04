package routes

import (
	"wallet-topup/internal/api/wallet/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *handlers.WalletHandler) {
	wallet := r.Group("/wallet")
	{
		wallet.POST("/verify", handler.VerifyTransaction)
		wallet.POST("/confirm/:transaction_id", handler.ConfirmTransaction)
	}
}
