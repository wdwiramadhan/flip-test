package service

import (
	"flip-test/internal/domain"
	"flip-test/internal/repository"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestSaveTransactions_ValidTransactions(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "John Doe",
			Type:   domain.TransactionTypeCredit,
			Amount: 1000000,
			Status: domain.TransactionStatusSuccess,
		},
	}

	err := service.SaveTransactions(transactions)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	saved := repo.GetTransactions()
	if len(saved) != 1 {
		t.Errorf("Expected 1 transaction to be saved, got %d", len(saved))
	}
}

func TestSaveTransactions_InvalidAmount(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "John Doe",
			Amount: 0, // Invalid: zero amount
			Status: domain.TransactionStatusSuccess,
		},
	}

	err := service.SaveTransactions(transactions)

	if err == nil {
		t.Error("Expected error for zero amount, got none")
	}

	if !strings.Contains(err.Error(), "invalid amount") {
		t.Errorf("Expected error message about invalid amount, got: %v", err)
	}
}

func TestSaveTransactions_NegativeAmount(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "John Doe",
			Amount: -1000, // Invalid: negative amount
			Status: domain.TransactionStatusSuccess,
		},
	}

	err := service.SaveTransactions(transactions)

	if err == nil {
		t.Error("Expected error for negative amount, got none")
	}
}

func TestSaveTransactions_EmptyName(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "", // Invalid: empty name
			Amount: 1000000,
			Status: domain.TransactionStatusSuccess,
		},
	}

	err := service.SaveTransactions(transactions)

	if err == nil {
		t.Error("Expected error for empty name, got none")
	}

	if !strings.Contains(err.Error(), "invalid name") {
		t.Errorf("Expected error message about invalid name, got: %v", err)
	}
}

func TestSaveTransactions_WhitespaceName(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "   ", // Invalid: whitespace only
			Amount: 1000000,
			Status: domain.TransactionStatusSuccess,
		},
	}

	err := service.SaveTransactions(transactions)

	if err == nil {
		t.Error("Expected error for whitespace-only name, got none")
	}
}

func TestGetBalance_EmptyTransactions(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	balance := service.GetBalance()

	if balance != 0 {
		t.Errorf("Expected balance 0, got %d", balance)
	}
}

func TestGetBalance_OnlyCreditTransactions(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "Transaction 1",
			Type:   domain.TransactionTypeCredit,
			Amount: 1000000,
			Status: domain.TransactionStatusSuccess,
		},
		{
			ID:     uuid.New(),
			Name:   "Transaction 2",
			Type:   domain.TransactionTypeCredit,
			Amount: 500000,
			Status: domain.TransactionStatusSuccess,
		},
	}

	repo.SaveTransactions(transactions)
	balance := service.GetBalance()

	expected := int64(1500000)
	if balance != expected {
		t.Errorf("Expected balance %d, got %d", expected, balance)
	}
}

func TestGetBalance_OnlyDebitTransactions(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "Transaction 1",
			Type:   domain.TransactionTypeDebit,
			Amount: 300000,
			Status: domain.TransactionStatusSuccess,
		},
		{
			ID:     uuid.New(),
			Name:   "Transaction 2",
			Type:   domain.TransactionTypeDebit,
			Amount: 200000,
			Status: domain.TransactionStatusSuccess,
		},
	}

	repo.SaveTransactions(transactions)
	balance := service.GetBalance()

	expected := int64(-500000)
	if balance != expected {
		t.Errorf("Expected balance %d, got %d", expected, balance)
	}
}

func TestGetBalance_MixedTransactions(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "Credit",
			Type:   domain.TransactionTypeCredit,
			Amount: 1000000,
			Status: domain.TransactionStatusSuccess,
		},
		{
			ID:     uuid.New(),
			Name:   "Debit",
			Type:   domain.TransactionTypeDebit,
			Amount: 300000,
			Status: domain.TransactionStatusSuccess,
		},
	}

	repo.SaveTransactions(transactions)
	balance := service.GetBalance()

	expected := int64(700000) // 1000000 - 300000
	if balance != expected {
		t.Errorf("Expected balance %d, got %d", expected, balance)
	}
}

func TestGetBalance_IgnoresFailedTransactions(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "Successful Credit",
			Type:   domain.TransactionTypeCredit,
			Amount: 1000000,
			Status: domain.TransactionStatusSuccess,
		},
		{
			ID:     uuid.New(),
			Name:   "Failed Debit",
			Type:   domain.TransactionTypeDebit,
			Amount: 500000,
			Status: domain.TransactionStatusFailed,
		},
	}

	repo.SaveTransactions(transactions)
	balance := service.GetBalance()

	expected := int64(1000000) // Failed transaction should not affect balance
	if balance != expected {
		t.Errorf("Expected balance %d, got %d", expected, balance)
	}
}

func TestGetBalance_IgnoresPendingTransactions(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "Successful Credit",
			Type:   domain.TransactionTypeCredit,
			Amount: 1000000,
			Status: domain.TransactionStatusSuccess,
		},
		{
			ID:     uuid.New(),
			Name:   "Pending Debit",
			Type:   domain.TransactionTypeDebit,
			Amount: 300000,
			Status: domain.TransactionStatusPending,
		},
	}

	repo.SaveTransactions(transactions)
	balance := service.GetBalance()

	expected := int64(1000000) // Pending transaction should not affect balance
	if balance != expected {
		t.Errorf("Expected balance %d, got %d", expected, balance)
	}
}

func TestGetUnsuccessfulTransactions_EmptyRepository(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	result := service.GetUnsuccessfulTransactions()

	if len(result) != 0 {
		t.Errorf("Expected 0 unsuccessful transactions, got %d", len(result))
	}
}

func TestGetUnsuccessfulTransactions_OnlySuccessful(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "Transaction 1",
			Status: domain.TransactionStatusSuccess,
		},
		{
			ID:     uuid.New(),
			Name:   "Transaction 2",
			Status: domain.TransactionStatusSuccess,
		},
	}

	repo.SaveTransactions(transactions)
	result := service.GetUnsuccessfulTransactions()

	if len(result) != 0 {
		t.Errorf("Expected 0 unsuccessful transactions, got %d", len(result))
	}
}

func TestGetUnsuccessfulTransactions_ReturnsFailedAndPending(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	transactions := []domain.Transaction{
		{
			ID:     uuid.New(),
			Name:   "Successful",
			Status: domain.TransactionStatusSuccess,
		},
		{
			ID:     uuid.New(),
			Name:   "Failed",
			Status: domain.TransactionStatusFailed,
		},
		{
			ID:     uuid.New(),
			Name:   "Pending",
			Status: domain.TransactionStatusPending,
		},
	}

	repo.SaveTransactions(transactions)
	result := service.GetUnsuccessfulTransactions()

	if len(result) != 2 {
		t.Fatalf("Expected 2 unsuccessful transactions, got %d", len(result))
	}

	// Check that we have failed and pending, but not successful
	statuses := make(map[domain.TransactionStatus]bool)
	for _, tx := range result {
		statuses[tx.Status] = true
	}

	if statuses[domain.TransactionStatusSuccess] {
		t.Error("Should not include successful transactions")
	}
	if !statuses[domain.TransactionStatusFailed] {
		t.Error("Should include failed transactions")
	}
	if !statuses[domain.TransactionStatusPending] {
		t.Error("Should include pending transactions")
	}
}

func TestGetUnsuccessfulTransactions_SortedByDateDescending(t *testing.T) {
	repo := repository.NewTransactionRepository()
	service := NewTransactionService(repo)

	now := time.Now()
	transactions := []domain.Transaction{
		{
			ID:              uuid.New(),
			Name:            "Oldest",
			Status:          domain.TransactionStatusFailed,
			TransactionDate: now.Add(-2 * time.Hour),
		},
		{
			ID:              uuid.New(),
			Name:            "Newest",
			Status:          domain.TransactionStatusFailed,
			TransactionDate: now,
		},
		{
			ID:              uuid.New(),
			Name:            "Middle",
			Status:          domain.TransactionStatusPending,
			TransactionDate: now.Add(-1 * time.Hour),
		},
	}

	repo.SaveTransactions(transactions)
	result := service.GetUnsuccessfulTransactions()

	if len(result) != 3 {
		t.Fatalf("Expected 3 unsuccessful transactions, got %d", len(result))
	}

	// Check sorting: should be descending (newest first)
	if result[0].Name != "Newest" {
		t.Errorf("Expected first transaction to be 'Newest', got '%s'", result[0].Name)
	}
	if result[1].Name != "Middle" {
		t.Errorf("Expected second transaction to be 'Middle', got '%s'", result[1].Name)
	}
	if result[2].Name != "Oldest" {
		t.Errorf("Expected third transaction to be 'Oldest', got '%s'", result[2].Name)
	}

	// Verify dates are in descending order
	for i := 0; i < len(result)-1; i++ {
		if result[i].TransactionDate.Before(result[i+1].TransactionDate) {
			t.Error("Transactions are not sorted in descending order by date")
		}
	}
}
