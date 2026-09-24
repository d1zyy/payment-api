package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/d1zyy/payment-api/internal/model"
)

func (h *PaymentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid payment", http.StatusBadRequest)
		return
	}

	payment, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, model.ErrPaymentNotFound) {
			http.Error(w, "payment not found", http.StatusNotFound)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := PaymentResponse{
		ID:     payment.ID,
		Amount: payment.Amount,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}

}
