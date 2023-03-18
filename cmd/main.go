// Package main - Read data from a libre sensor api
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JOHNEPPILLAR/flowercare-data-collector/internal/fcServer"
	"github.com/JOHNEPPILLAR/utility/database"
	"github.com/JOHNEPPILLAR/utility/logger"
)

func main() {

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initialize Logger
	logger, logErr := logger.NewLogger()
	if logErr != nil {
		logger.Fatal("Unable to create logger")
		os.Exit(1)
	}

	logger.Info("Starting Flowercare Data Collector Server...")

	// Initialize Database connection
	db, dbErr := database.NewMongoConnection(ctx, logger).ConnectDB()
	if dbErr != nil {
		logger.Fatal("Unable to initialize database")
		os.Exit(1)
	}
	defer db.Disconnect(ctx)

	logger.Info("Database online")

	// Start flowercare data collector server
	go func() {
		fcServer.CreateServer(ctx, logger, db).Start()
	}()

	// Listen for interrupt signal
	<-ctx.Done()

	// Restore default behavior on the interrupt signal
	stop()
	logger.Warn("Received interrupt signal, shutting down...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("Shutdown")
}
