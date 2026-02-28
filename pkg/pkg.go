package pkg

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/MediStatTech/dashboard-service/pkg/config"
	"github.com/MediStatTech/logger"
)

type Facade struct {
	Logger    *logger.Logger
	Config    *config.Config
}

func New(ctx context.Context) (*Facade, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create config: %w", err)
	}

	logger := initLogger()

	return &Facade{
		Logger:    logger,
		Config:    config,
	}, nil
}

func initLogger() *logger.Logger {
	// TODO: Implement logger configuration
	return logger.New(io.Writer(os.Stdout))
}
