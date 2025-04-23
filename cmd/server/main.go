package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pvnptl/task-service/internal/api"
	"github.com/pvnptl/task-service/internal/api/handlers"
	"github.com/pvnptl/task-service/internal/repository"
	"github.com/pvnptl/task-service/internal/service"
)

func main() {
	// Setup everything needed
	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)
	router := api.NewRouter(taskHandler)

	// Create the server
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	// Start the server in the background
	go func() {
		log.Println("Server started at http://localhost:8081")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()


	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Shutdown the server
	log.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	log.Println("Server stopped")
}