package handler

import (
	"encoding/json"
	"net/http"
)

type CreatePaymentRequest struct {
	Amount int `json:"amount"`
}

type CreatePaymentResponse struct {
	ID     int `json:"id"`
	Amount int `json:"amount"`
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var req CreatePaymentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "amount must be greater than 0", http.StatusUnprocessableEntity)
		return
	}

	payment, err := h.service.Create(r.Context(), req.Amount)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	response := CreatePaymentResponse{
		ID:     payment.ID,
		Amount: payment.Amount,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
