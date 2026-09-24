package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/d1zyy/payment-api/internal/handler"
	"github.com/d1zyy/payment-api/internal/repository"
	"github.com/d1zyy/payment-api/internal/service"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//service payment
	paymentRepo := repository.NewMemoryPaymentRepository()
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	//endpoint
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("POST /payments", paymentHandler.CreatePayment)
	mux.HandleFunc("GET /get/{id}", paymentHandler.GetByID)

	//graceful shutdown
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("server start on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error start server: %v", err)
		}
	}()

	<-shutdownChan
	log.Println("stop server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("critical stop server: %v", err)
	}

	log.Println("server shutdown successfully")
}
