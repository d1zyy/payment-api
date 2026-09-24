package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/d1zyy/payment-api/internal/model"
)

func TestMemoryPayment_Unique(t *testing.T) {
	repo := NewMemoryPaymentRepository()
	ctx := context.Background()

	payment1, err := repo.Create(ctx, model.Payment{
		Amount: 100,
	})
	if err != nil {
		t.Fatalf("failed to create first payment: %v", err)
	}

	payment2, err := repo.Create(ctx, model.Payment{
		Amount: 200,
	})
	if err != nil {
		t.Fatalf("failed to create second payment: %v", err)
	}

	if payment1.ID == payment2.ID {
		t.Fatalf("expected unique ID, got %d and %d", payment1.ID, payment2.ID)
	}
}

func TestMemoryIDtoAmount(t *testing.T) {
	repo := NewMemoryPaymentRepository()
	ctx := context.Background()

	expected, err := repo.Create(ctx, model.Payment{
		Amount: 100,
	})
	if err != nil {
		t.Fatalf("failed to create payment: %v", err)
	}

	actual, err := repo.GetByID(ctx, expected.ID)
	if err != nil {
		t.Fatalf("failed to get payment: %v", err)
	}

	if actual.ID != expected.ID {
		t.Fatalf("expected ID %d, got %d", expected.ID, actual.ID)
	}

	if actual.Amount != expected.Amount {
		t.Fatalf("expected amount %d, got %d", expected.Amount, actual.Amount)
	}

}

func TestMemoryPaymentRepo_GetByID_NotFound(t *testing.T) {
	repo := NewMemoryPaymentRepository()
	ctx := context.Background()

	_, err := repo.GetByID(ctx, 999)
	if err == nil {
		t.Fatalf("failed to get ID: %v", err)
	}

	if !errors.Is(err, model.ErrPaymentNotFound) {
		t.Fatalf("expected ErrPaymentNotFound, got %v", err)
	}
}
