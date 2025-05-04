package handlers

import (
	"net/http"
	"wallet-topup/internal/api/wallet"
	"wallet-topup/internal/global/responses"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletHandler struct {
	service wallet.Service
}

func NewHandler(service wallet.Service) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) VerifyTransaction(c *gin.Context) {
	var req wallet.VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.service.VerifyTransaction(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}

func (h *WalletHandler) ConfirmTransaction(c *gin.Context) {
	transactionIDStr := c.Param("transaction_id")
	if transactionIDStr == "" {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest{
			Code:    "400",
			Message: "transaction_id is required",
			Data:    nil,
		})
		return
	}

	transactionID, err := uuid.Parse(transactionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest{
			Code:    "400",
			Message: "invalid transaction_id",
			Data:    nil,
		})
		return
	}

	tx, err := h.service.ConfirmTransaction(c.Request.Context(), transactionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest{
			Code:    "400",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, responses.ResponseSuccess{
		Code:    "200",
		Message: "success",
		Data:    tx,
	})
}
