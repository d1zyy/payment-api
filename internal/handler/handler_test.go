package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/d1zyy/payment-api/internal/handler"
	"github.com/d1zyy/payment-api/internal/model"
	"github.com/d1zyy/payment-api/internal/repository"
	"github.com/d1zyy/payment-api/internal/service"
)

func TestPaymentHandler_GetByID_InvalidID(t *testing.T) {
	repo := repository.NewMemoryPaymentRepository()
	svc := service.NewPaymentService(repo)
	h := handler.NewPaymentHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /payments/{id}", h.GetByID)

	req := httptest.NewRequest(
		http.MethodGet,
		"/payments/abc",
		nil,
	)

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestPaymentHandler_GetByID_NotFound(t *testing.T) {
	repo := repository.NewMemoryPaymentRepository()
	svc := service.NewPaymentService(repo)
	h := handler.NewPaymentHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /payments/{id}", h.GetByID)

	req := httptest.NewRequest(
		http.MethodGet,
		"/payments/999",
		nil,
	)

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestPaymentHandler_GetByID_OK(t *testing.T) {
	repo := repository.NewMemoryPaymentRepository()
	svc := service.NewPaymentService(repo)
	h := handler.NewPaymentHandler(svc)

	ctx := context.Background()

	created, err := repo.Create(ctx, model.Payment{
		Amount: 800,
	})
	if err != nil {
		t.Fatalf("failed to create payment: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /payments/{id}", h.GetByID)

	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/payments/%d", created.ID),
		nil,
	)

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	expectedJSON := `{
		"id": 1,
		"amount":    800
	}`

	require.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	assert.JSONEq(t, expectedJSON, rec.Body.String())

}
