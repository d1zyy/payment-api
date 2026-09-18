package service

import (
	"sync"
)

type Payment struct {
	ID     int
	Amount int
}

type PaymentService struct {
	mu     sync.Mutex
	nextID int
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		nextID: 1,
	}
}

func (s *PaymentService) Create(amount int) Payment {
	s.mu.Lock()
	defer s.mu.Unlock()

	payment := Payment{
		ID:     s.nextID,
		Amount: amount,
	}

	s.nextID++

	return payment
}
