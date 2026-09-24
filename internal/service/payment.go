package service

import (
	"context"

	"github.com/d1zyy/payment-api/internal/model"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment model.Payment) (model.Payment, error)
	GetByID(ctx context.Context, id int) (model.Payment, error)
}

type PaymentService struct {
	repo PaymentRepository
}

func NewPaymentService(repo PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (s *PaymentService) Create(ctx context.Context, amount int) (model.Payment, error) {
	payment := model.Payment{
		Amount: amount,
	}

	return s.repo.Create(ctx, payment)
}

func (s *PaymentService) GetByID(ctx context.Context, id int) (model.Payment, error) {
	return s.repo.GetByID(ctx, id)
}
