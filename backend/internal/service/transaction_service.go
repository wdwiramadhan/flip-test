package service

import (
	"flip-test/internal/domain"
	"flip-test/internal/repository"
	"fmt"
	"sort"
	"strings"
)

type TransactionService struct {
	TransactionRepository *repository.TransactionRepository
}

func NewTransactionService(tr *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		TransactionRepository: tr,
	}
}

func (ts *TransactionService) SaveTransactions(transactions []domain.Transaction) error {
	for i, transaction := range transactions {
		if transaction.Amount <= 0 {
			return fmt.Errorf("invalid amount at row %d: amount must be greater than 0", i+1)
		}

		if strings.TrimSpace(transaction.Name) == "" {
			return fmt.Errorf("invalid name at row %d: name cannot be empty", i+1)
		}
	}

	ts.TransactionRepository.SaveTransactions(transactions)
	return nil
}

func (ts TransactionService) GetBalance() int64 {
	var balance int64 = 0
	transactions := ts.TransactionRepository.GetTransactions()

	for _, transaction := range transactions {
		if transaction.Status != domain.TransactionStatusSuccess {
			continue
		}

		if transaction.Type == domain.TransactionTypeCredit {
			balance += transaction.Amount
		} else {
			balance -= transaction.Amount
		}
	}

	return balance
}

func (ts TransactionService) GetUnsuccessfulTransactions() []domain.Transaction {
	failedTransactions := make([]domain.Transaction, 0)
	transactions := ts.TransactionRepository.GetTransactions()
	for _, transaction := range transactions {
		if transaction.Status != domain.TransactionStatusSuccess {
			failedTransactions = append(failedTransactions, transaction)
		}
	}

	sort.Slice(failedTransactions, func(i, j int) bool {
		return failedTransactions[i].TransactionDate.After(failedTransactions[j].TransactionDate)
	})

	return failedTransactions
}
