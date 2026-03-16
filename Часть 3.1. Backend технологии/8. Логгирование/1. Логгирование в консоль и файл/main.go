package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DEBUG
// INFO
// WARN
// ERROR

func NewLogger(logLevel string) (*zap.Logger, func() error, error) {
	lvl := zap.NewAtomicLevel()
	if err := lvl.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, nil, fmt.Errorf("unmarshal log level: %w", err)
	}

	// 7           5              5
	// user        user-group     other
	// r w x       r w x          r w x
	// 1 1 1       1 0 1          1 0 1
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, nil, fmt.Errorf("mkdir log folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join("logs", fmt.Sprintf("%s.log", timestamp))

	// 6           4              4
	// user        user-group     other
	// r w x       r w x          r w x
	// 1 1 0       1 0 0          1 0 0
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("open log file: %w", err)
	}

	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")

	encoder := zapcore.NewConsoleEncoder(cfg)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lvl),
		zapcore.NewCore(encoder, zapcore.AddSync(logFile), lvl),
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return logger, logFile.Close, nil
}

func main() {
	logger, logFileClose, err := NewLogger("DEBUG")
	if err != nil {
		panic(err)
	}
	defer logFileClose()

	var (
		errTest = errors.New("TEST ERROR")
	)

	logger.Debug("Debug logging", zap.Time("time_field", time.Now().UTC()))
	logger.Info("Some Info log", zap.Int("int_field", 3))
	logger.Warn("Warning logging", zap.Float64("float_field", 5.25))
	logger.Error("Some Error logging", zap.Error(errTest))
}
