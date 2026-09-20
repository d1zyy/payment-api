package repository

import (
	"context"
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
