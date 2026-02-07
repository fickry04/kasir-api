package services

import (
	"kasir-api/internal/models"
	"kasir-api/internal/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) Summary(startDate string, endDate string) (*models.Summary, error) {
	return s.repo.Summary(startDate, endDate)
}

func (s *TransactionService) SummaryToday() (*models.Summary, error) {
	return s.repo.SummaryToday()
}
