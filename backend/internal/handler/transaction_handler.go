package handler

import (
	"flip-test/internal/parser"
	"flip-test/internal/service"
	"log"
	"net/http"
	"strings"
)

type TransactionHandler struct {
	TransactionService *service.TransactionService
}

func NewTransactionHandler(ts *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: ts,
	}
}

func (th *TransactionHandler) UploadCSV(w http.ResponseWriter, req *http.Request) {
	// Parse the multipart form with 10 MB memory limit
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Failed to parse multipart form: %v", err)
		WriteJSON(w, http.StatusBadRequest, "BAD_REQUEST", "Failed to parse form", nil)
		return
	}

	file, header, err := req.FormFile("file")
	if err != nil {
		log.Printf("Failed to get file from form: %v", err)
		WriteJSON(w, http.StatusBadRequest, "BAD_REQUEST", "CSV file is required", nil)
		return
	}
	defer file.Close()

	if !strings.HasSuffix(strings.ToLower(header.Filename), ".csv") {
		log.Printf("Invalid file extension: %s", header.Filename)
		WriteJSON(w, http.StatusBadRequest, "BAD_REQUEST", "Only CSV files are allowed", nil)
		return
	}

	if header.Size > 10<<20 {
		log.Printf("File too large: %d bytes", header.Size)
		WriteJSON(w, http.StatusBadRequest, "BAD_REQUEST", "File size must be less than 10MB", nil)
		return
	}

	log.Printf("Processing CSV file: %s (size: %d bytes)", header.Filename, header.Size)

	transactions, err := parser.ParseCSVToTransactions(file)
	if err != nil {
		log.Printf("Failed to parse CSV: %v", err)
		WriteJSON(w, http.StatusBadRequest, "BAD_REQUEST", err.Error(), nil)
		return
	}

	log.Printf("Parsed %d transactions from CSV", len(transactions))

	err = th.TransactionService.SaveTransactions(transactions)
	if err != nil {
		log.Printf("Failed to save transactions: %v", err)
		WriteJSON(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to save transactions", nil)
		return
	}

	log.Printf("Successfully saved %d transactions", len(transactions))
	WriteJSON(w, http.StatusOK, "SUCCESS", "Transactions uploaded successfully", nil)
}

func (th *TransactionHandler) GetBalance(w http.ResponseWriter, req *http.Request) {
	balance := th.TransactionService.GetBalance()
	WriteJSON(w, http.StatusOK, "SUCCESS", "SUCCESS", balance)
}

func (th *TransactionHandler) GetUnsuccessfulTransactions(w http.ResponseWriter, req *http.Request) {
	transactions := th.TransactionService.GetUnsuccessfulTransactions()
	WriteJSON(w, http.StatusOK, "SUCCESS", "SUCCESS", transactions)
}
