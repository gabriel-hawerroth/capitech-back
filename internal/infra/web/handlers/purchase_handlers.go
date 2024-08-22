package handlers

import "github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"

type PurchaseHandler struct {
	PurchaseRepository repositories.PurchaseRepository
}

func NewPurchaseHandler(purchaseRepository repositories.PurchaseRepository) *PurchaseHandler {
	return &PurchaseHandler{
		PurchaseRepository: purchaseRepository,
	}
}
