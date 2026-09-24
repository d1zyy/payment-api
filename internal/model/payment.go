package model

import "errors"

var ErrPaymentNotFound = errors.New("payment not found")

type Payment struct {
	ID     int
	Amount int
}
