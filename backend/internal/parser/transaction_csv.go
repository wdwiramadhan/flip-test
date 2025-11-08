package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"flip-test/internal/domain"

	"github.com/google/uuid"
)

var expectedHeaders = []string{"timestamp", "name", "type", "amount", "status", "description"}

func ParseCSVToTransactions(r io.Reader) ([]domain.Transaction, error) {
	reader := csv.NewReader(r)

	if err := validateCSVHeaders(reader); err != nil {
		return nil, err
	}

	var transactions []domain.Transaction
	lineNum := 1 // Header is line 1, data starts at line 2
	for {
		lineNum++
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("line %d: failed to read row: %w", lineNum, err)
		}

		transaction, err := parseTransactionRow(record, lineNum)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func validateCSVHeaders(reader *csv.Reader) error {
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	if len(headers) != len(expectedHeaders) {
		return fmt.Errorf("invalid header count: expected %d columns, got %d", len(expectedHeaders), len(headers))
	}

	for i, header := range headers {
		normalizedHeader := strings.ToLower(strings.TrimSpace(header))
		if normalizedHeader != expectedHeaders[i] {
			return fmt.Errorf("invalid header at column %d: expected '%s', got '%s'", i+1, expectedHeaders[i], header)
		}
	}

	return nil
}

func parseTransactionRow(record []string, lineNum int) (domain.Transaction, error) {
	if len(record) != len(expectedHeaders) {
		return domain.Transaction{}, fmt.Errorf("line %d: expected %d columns, got %d", lineNum, len(expectedHeaders), len(record))
	}

	timestamp, err := strconv.ParseInt(strings.TrimSpace(record[0]), 10, 64)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("line %d: invalid timestamp '%s': %w", lineNum, record[0], err)
	}
	transactionDate := time.Unix(timestamp, 0).UTC()

	transactionType := domain.TransactionType(strings.TrimSpace(record[2]))
	if err := validateTransactionType(transactionType, lineNum, record[2]); err != nil {
		return domain.Transaction{}, err
	}

	amount, err := strconv.ParseInt(strings.TrimSpace(record[3]), 10, 64)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("line %d: invalid amount '%s': %w", lineNum, record[3], err)
	}

	transactionStatus := domain.TransactionStatus(strings.TrimSpace(record[4]))
	if err := validateTransactionStatus(transactionStatus, lineNum, record[4]); err != nil {
		return domain.Transaction{}, err
	}

	return domain.Transaction{
		ID:              uuid.New(),
		Name:            strings.TrimSpace(record[1]),
		Type:            transactionType,
		Amount:          amount,
		Status:          transactionStatus,
		Description:     strings.TrimSpace(record[5]),
		TransactionDate: transactionDate,
	}, nil
}

func validateTransactionType(t domain.TransactionType, lineNum int, original string) error {
	if t != domain.TransactionTypeDebit && t != domain.TransactionTypeCredit {
		return fmt.Errorf("line %d: invalid transaction type '%s'. Must be 'DEBIT' or 'CREDIT'", lineNum, original)
	}
	return nil
}

func validateTransactionStatus(s domain.TransactionStatus, lineNum int, original string) error {
	if s != domain.TransactionStatusSuccess &&
		s != domain.TransactionStatusPending &&
		s != domain.TransactionStatusFailed {
		return fmt.Errorf("line %d: invalid status '%s'. Must be 'SUCCESS', 'PENDING', or 'FAILED'", lineNum, original)
	}
	return nil
}
