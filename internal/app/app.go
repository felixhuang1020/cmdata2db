package app

import (
	"cmdata2db/internal/middleware"
	"fmt"

	"go.uber.org/zap"
)

// InitializeAll 初始化所有模块
func InitializeAll() error {
	// 初始化MySQL
	err := InitializeCk()
	if err != nil {
		// 使用 zap logger 记录错误
		middleware.GetLogger().Error("Clickhouse初始化错误", zap.Error(err))
		return fmt.Errorf("clickhouse初始化错误: %v", err)
	}

	return nil
}
