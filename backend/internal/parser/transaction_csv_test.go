package parser

import (
	"flip-test/internal/domain"
	"strings"
	"testing"
)

func TestParseCSVToTransactions_ValidCSV(t *testing.T) {
	csvData := `timestamp,name,type,amount,status,description
1704067200,John Doe,CREDIT,1000000,SUCCESS,Initial deposit
1704153600,Jane Smith,DEBIT,250000,FAILED,Failed payment`

	reader := strings.NewReader(csvData)
	transactions, err := ParseCSVToTransactions(reader)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(transactions) != 2 {
		t.Fatalf("Expected 2 transactions, got: %d", len(transactions))
	}

	// Test first transaction
	first := transactions[0]
	if first.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got: %s", first.Name)
	}
	if first.Type != domain.TransactionTypeCredit {
		t.Errorf("Expected type CREDIT, got: %s", first.Type)
	}
	if first.Amount != 1000000 {
		t.Errorf("Expected amount 1000000, got: %d", first.Amount)
	}
	if first.Status != domain.TransactionStatusSuccess {
		t.Errorf("Expected status SUCCESS, got: %s", first.Status)
	}

	// Test second transaction
	second := transactions[1]
	if second.Type != domain.TransactionTypeDebit {
		t.Errorf("Expected type DEBIT, got: %s", second.Type)
	}
	if second.Status != domain.TransactionStatusFailed {
		t.Errorf("Expected status FAILED, got: %s", second.Status)
	}
}

func TestParseCSVToTransactions_InvalidHeaders(t *testing.T) {
	csvData := `invalid,headers,here
1704067200,John Doe,CREDIT,1000000,SUCCESS,Initial deposit`

	reader := strings.NewReader(csvData)
	_, err := ParseCSVToTransactions(reader)

	if err == nil {
		t.Fatal("Expected error for invalid headers, got none")
	}
}

func TestParseCSVToTransactions_MissingHeaders(t *testing.T) {
	csvData := `timestamp,name,type,amount,status
1704067200,John Doe,CREDIT,1000000,SUCCESS`

	reader := strings.NewReader(csvData)
	_, err := ParseCSVToTransactions(reader)

	if err == nil {
		t.Fatal("Expected error for missing headers, got none")
	}
}

func TestParseCSVToTransactions_InvalidTransactionType(t *testing.T) {
	csvData := `timestamp,name,type,amount,status,description
1704067200,John Doe,INVALID,1000000,SUCCESS,Initial deposit`

	reader := strings.NewReader(csvData)
	_, err := ParseCSVToTransactions(reader)

	if err == nil {
		t.Fatal("Expected error for invalid transaction type, got none")
	}

	if !strings.Contains(err.Error(), "invalid transaction type") {
		t.Errorf("Expected error message about invalid transaction type, got: %v", err)
	}
}

func TestParseCSVToTransactions_InvalidStatus(t *testing.T) {
	csvData := `timestamp,name,type,amount,status,description
1704067200,John Doe,CREDIT,1000000,INVALID,Initial deposit`

	reader := strings.NewReader(csvData)
	_, err := ParseCSVToTransactions(reader)

	if err == nil {
		t.Fatal("Expected error for invalid status, got none")
	}

	if !strings.Contains(err.Error(), "invalid status") {
		t.Errorf("Expected error message about invalid status, got: %v", err)
	}
}

func TestParseCSVToTransactions_InvalidAmount(t *testing.T) {
	csvData := `timestamp,name,type,amount,status,description
1704067200,John Doe,CREDIT,not_a_number,SUCCESS,Initial deposit`

	reader := strings.NewReader(csvData)
	_, err := ParseCSVToTransactions(reader)

	if err == nil {
		t.Fatal("Expected error for invalid amount, got none")
	}

	if !strings.Contains(err.Error(), "invalid amount") {
		t.Errorf("Expected error message about invalid amount, got: %v", err)
	}
}

func TestParseCSVToTransactions_InvalidTimestamp(t *testing.T) {
	csvData := `timestamp,name,type,amount,status,description
invalid_timestamp,John Doe,CREDIT,1000000,SUCCESS,Initial deposit`

	reader := strings.NewReader(csvData)
	_, err := ParseCSVToTransactions(reader)

	if err == nil {
		t.Fatal("Expected error for invalid timestamp, got none")
	}

	if !strings.Contains(err.Error(), "invalid timestamp") {
		t.Errorf("Expected error message about invalid timestamp, got: %v", err)
	}
}

func TestParseCSVToTransactions_EmptyCSV(t *testing.T) {
	csvData := `timestamp,name,type,amount,status,description`

	reader := strings.NewReader(csvData)
	transactions, err := ParseCSVToTransactions(reader)

	if err != nil {
		t.Fatalf("Expected no error for empty CSV, got: %v", err)
	}

	if len(transactions) != 0 {
		t.Errorf("Expected 0 transactions, got: %d", len(transactions))
	}
}

func TestParseCSVToTransactions_WithWhitespace(t *testing.T) {
	csvData := `timestamp,name,type,amount,status,description
1704067200,  John Doe  ,  CREDIT  ,  1000000  ,  SUCCESS  ,  Initial deposit  `

	reader := strings.NewReader(csvData)
	transactions, err := ParseCSVToTransactions(reader)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(transactions) != 1 {
		t.Fatalf("Expected 1 transaction, got: %d", len(transactions))
	}

	transaction := transactions[0]
	if transaction.Name != "John Doe" {
		t.Errorf("Expected trimmed name 'John Doe', got: '%s'", transaction.Name)
	}
	if transaction.Description != "Initial deposit" {
		t.Errorf("Expected trimmed description 'Initial deposit', got: '%s'", transaction.Description)
	}
}

func TestParseCSVToTransactions_WrongColumnCount(t *testing.T) {
	csvData := `timestamp,name,type,amount,status,description
1704067200,John Doe,CREDIT,1000000`

	reader := strings.NewReader(csvData)
	_, err := ParseCSVToTransactions(reader)

	if err == nil {
		t.Fatal("Expected error for wrong column count, got none")
	}

	if !strings.Contains(err.Error(), "wrong number of fields") {
		t.Errorf("Expected error message about wrong number of fields, got: %v", err)
	}
}
