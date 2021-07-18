package sample

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Version metadata set by ldflags during the build.
var (
	version string
	commit  string
	date    string
)

// Version returns a string with version metadata: version number, git sha and build date.
// It returns "development" if version variables are not set during the build.
func Version() string {
	if version == "" {
		return "development"
	}

	return fmt.Sprintf("%s - revision %s built at %s", version, commit[:6], date)
}

// Logger return a zap.SugaredLogger configured with a specified log level.
func Logger(level string) (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.DisableStacktrace = true

	switch strings.ToLower(level) {
	case "debug":
		cfg.Level.SetLevel(zap.DebugLevel)
	case "info":
		cfg.Level.SetLevel(zap.InfoLevel)
	case "error":
		cfg.Level.SetLevel(zap.ErrorLevel)
	default:
		return nil, fmt.Errorf("invalid log level")
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	sugar := logger.Sugar()

	return sugar, nil
}
