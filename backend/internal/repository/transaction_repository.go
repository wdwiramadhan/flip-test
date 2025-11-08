package repository

import (
	"flip-test/internal/domain"
	"sync"

	"github.com/google/uuid"
)

type TransactionRepository struct {
	store map[uuid.UUID]domain.Transaction
	mutex sync.RWMutex
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{store: make(map[uuid.UUID]domain.Transaction)}
}

func (tr *TransactionRepository) GetTransactions() []domain.Transaction {
	tr.mutex.RLock()
	defer tr.mutex.RUnlock()

	result := make([]domain.Transaction, 0, len(tr.store))
	for _, transaction := range tr.store {
		result = append(result, transaction)
	}

	return result
}

func (tr *TransactionRepository) SaveTransactions(transactions []domain.Transaction) {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	for _, transaction := range transactions {
		tr.store[transaction.ID] = transaction
	}
}
