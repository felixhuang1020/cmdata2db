package utils

import (
	"strings"

	"go.uber.org/zap"
)

type LoggerWriter struct {
	Logger *zap.Logger
	Level  string // 日志级别："info" 或 "error"
}

func (w *LoggerWriter) Write(p []byte) (n int, err error) {
	// 去除日志末尾的换行符
	msg := strings.TrimSpace(string(p))
	if msg == "" {
		return len(p), nil
	}

	// 根据级别记录日志
	switch w.Level {
	case "info":
		w.Logger.Info(msg)
	case "error":
		w.Logger.Error(msg)
	default:
		w.Logger.Info(msg)
	}
	return len(p), nil
}
