package repository

import (
	"context"
	"sync"

	"github.com/d1zyy/payment-api/internal/model"
)

type MemoryPaymentRepository struct {
	mu       sync.RWMutex
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

func (r *MemoryPaymentRepository) GetByID(ctx context.Context, id int) (model.Payment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if payment, exists := r.payments[id]; exists {
		return payment, nil
	}

	return model.Payment{}, model.ErrPaymentNotFound
}
