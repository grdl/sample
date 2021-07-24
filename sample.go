package sample

import (
	"fmt"

	"go.uber.org/zap"
)

type Config struct {
	LogLevel string
}

func (c *Config) validate() error {
	validLogLevels := map[string]struct{}{
		"info":  {},
		"error": {},
		"debug": {},
	}

	if _, ok := validLogLevels[c.LogLevel]; !ok {
		return fmt.Errorf("level flag contains invalid value; valid values: %v", validLogLevels)
	}

	return nil
}

type Sample struct {
	config *Config
	logger *zap.SugaredLogger
}

func New(config *Config) (*Sample, error) {
	err := config.validate()
	if err != nil {
		return nil, err
	}

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
