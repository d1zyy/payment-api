package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/d1zyy/payment-api/internal/repository"
	"github.com/d1zyy/payment-api/internal/service"
)

func TestRepo(t *testing.T) {
	paymentRepo := repository.NewMemoryPaymentRepository()
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := NewPaymentHandler(paymentService)

	body := strings.NewReader(`{"amount":1000}`)

	req := httptest.NewRequest(http.MethodPost, "/payments", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	paymentHandler.CreatePayment(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

}
