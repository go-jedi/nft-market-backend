package service

import (
	"github.com/rob-bender/nft-market-backend/appl_row"
	"github.com/rob-bender/nft-market-backend/pkg/repository"
)

type PaymentService struct {
	repo repository.TodoPayment
}

func NewPaymentService(r repository.TodoPayment) *PaymentService {
	return &PaymentService{
		repo: r,
	}
}

func (s *PaymentService) CreatePayment(paymentForm appl_row.PaymentCreate) (int, error) {
	return s.repo.CreatePayment(paymentForm)
}

func (s *PaymentService) GetAllPayments() ([]appl_row.Payment, int, error) {
	return s.repo.GetAllPayments()
}
