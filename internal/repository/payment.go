package repository

import (
	"context"
	"sync"

	"github.com/d1zyy/payment-api/internal/model"
)

type MemoryPaymentRepository struct {
	mu       sync.Mutex
	nextID   int
	payments map[int]model.Payment
}

func NewMemoryPaymentRepository() *MemoryPaymentRepository {
	return &MemoryPaymentRepository{
		nextID:   1,
		payments: make(map[int]model.Payment),
	}
}

func (r *MemoryPaymentRepository) Create(ctx context.Context, payment model.Payment) (model.Payment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	payment.ID = r.nextID
	r.payments[payment.ID] = payment
	r.nextID++

	return payment, nil
}
