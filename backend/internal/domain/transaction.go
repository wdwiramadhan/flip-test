package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string
type TransactionStatus string

const (
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed  TransactionStatus = "FAILED"
	TransactionStatusPending TransactionStatus = "PENDING"
)

const (
	TransactionTypeDebit  TransactionType = "DEBIT"
	TransactionTypeCredit TransactionType = "CREDIT"
)

type Transaction struct {
	ID              uuid.UUID         `json:"id"`
	Name            string            `json:"name"`
	Type            TransactionType   `json:"type"`
	Amount          int64             `json:"amount"`
	Status          TransactionStatus `json:"status"`
	Description     string            `json:"description"`
	TransactionDate time.Time         `json:"transaction_date"`
}
