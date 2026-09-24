package handler

import "github.com/d1zyy/payment-api/internal/service"

type PaymentHandler struct {
	service *service.PaymentService
}

type PaymentResponse struct {
	ID     int `json:"id"`
	Amount int `json:"amount"`
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}
