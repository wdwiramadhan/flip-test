package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"flip-test/internal/handler"
	"flip-test/internal/middleware"
	"flip-test/internal/repository"
	"flip-test/internal/service"
)

func main() {
	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /transactions/upload", transactionHandler.UploadCSV)
	mux.HandleFunc("GET /transactions/balance", transactionHandler.GetBalance)
	mux.HandleFunc("GET /transactions/issues", transactionHandler.GetUnsuccessfulTransactions)

	handler := middleware.Chain(
		middleware.LoggingMiddleware,
		middleware.CorsMiddleware,
	)(mux)

	server := &http.Server{
		Addr:         getServerAddr(),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	gracefulShutdown(server)
}

func getServerAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
