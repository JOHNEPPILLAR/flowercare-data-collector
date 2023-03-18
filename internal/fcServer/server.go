// Package fcServer - Interact with flowercare BLE API's
package fcServer

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/JOHNEPPILLAR/utility/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// Start - Get and process data from flowercare devices
func (fcs *FlowerCare) Start() {

	fcs.logger.Info("Starting data collector server...")

	// Start health check api server
	go func() {
		fcs.healthCheckAPI()
	}()

	pollingTimer, err := strconv.Atoi(os.Getenv("POLLING_TIMER"))
	if err != nil {
		fcs.logger.Info("ENV: POLLING_TIMER missing, using default 30 minutes")
	}

	if pollingTimer == 0 {
		pollingTimer = 30
	}
	pollingInterval := time.Duration(pollingTimer) * time.Minute

	scanTicker := time.NewTicker(pollingInterval)

	// Run scan now
	fcs.scan()

	// Then every interval
	for range scanTicker.C {
		fcs.scan()
	}

	fcs.logger.Debug("Data collector server stopped")
}

// FlowerCare - Micro framework class
type FlowerCare struct {
	ctx        context.Context
	logger     *logger.Logger
	db         *mongo.Client
	httpClient *http.Client
	errorCount int
}

// CreateServer - New server instance
func CreateServer(ctx context.Context, logger *logger.Logger, db *mongo.Client) FCServer {
	return &FlowerCare{
		ctx:        ctx,
		logger:     logger,
		db:         db,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		errorCount: 0,
	}
}
