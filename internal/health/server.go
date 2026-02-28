package health

import (
	"context"
	"net/http"
	"time"

	"github.com/MediStatTech/logger"
)

// HealthServer provides HTTP endpoints for Kubernetes health checks
type HealthServer struct {
	log    *logger.Logger
	port   string
	server *http.Server
}

// NewHealthServer creates a new health check server
func NewHealthServer(log *logger.Logger, port string) *HealthServer {
	return &HealthServer{
		log:  log,
		port: port,
	}
}

// Start begins listening on the health port (non-blocking)
func (h *HealthServer) Start() error {
	mux := http.NewServeMux()

	// Liveness probe: Is the server alive?
	mux.HandleFunc("/healthz", h.handleHealthz)

	// Readiness probe: Is the server ready to accept traffic?
	mux.HandleFunc("/readyz", h.handleReadyz)

	h.server = &http.Server{
		Addr:              h.port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	h.log.Infof("Starting health server on %s", h.port)

	// Start server in background
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			h.log.Errorf("Health server error: %v", err)
		}
	}()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	return nil
}

// Shutdown gracefully stops the health server
func (h *HealthServer) Shutdown(ctx context.Context) error {
	if h.server == nil {
		return nil
	}
	h.log.Infof("Shutting down health server")
	return h.server.Shutdown(ctx)
}

// handleHealthz responds to liveness probes
func (h *HealthServer) handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// handleReadyz responds to readiness probes
func (h *HealthServer) handleReadyz(w http.ResponseWriter, r *http.Request) {
	// TODO: Add dependency checks here (DB, Redis, etc.)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("READY"))
}
