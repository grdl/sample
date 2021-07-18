package sample

import (
	"fmt"

	"go.uber.org/zap"
)

type Config struct {
	LogLevel string
}

type Sample struct {
	config *Config
	logger *zap.SugaredLogger
}

func New(config *Config) (*Sample, error) {
	logger, err := Logger(config.LogLevel)
	if err != nil {
		return nil, err
	}

	return &Sample{
		config: config,
		logger: logger,
	}, nil
}

func (s *Sample) Run() error {
	s.logger.Infow("test", "key", "info")
	s.logger.Errorw("test", "key", "error")
	s.logger.Debugw("test", "key", "debug")

	return fmt.Errorf("wot")
	//return nil
}
