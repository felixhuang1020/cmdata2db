package app

import (
	"cmdata2db/config"
	v1 "cmdata2db/internal/api/v1"
	"cmdata2db/internal/middleware"
	"cmdata2db/internal/utils"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Start() {
	var err error
	originalDefaultWriter := gin.DefaultWriter
	originalDefaultErrorWriter := gin.DefaultErrorWriter

	gin.DefaultWriter = io.MultiWriter(
		originalDefaultWriter, // 保留原始日志输出（如控制台）
		&utils.LoggerWriter{Logger: middleware.GetLogger(), Level: "info"}, // 同步到 zap info
	)

	gin.DefaultErrorWriter = io.MultiWriter(
		originalDefaultErrorWriter, // 保留原始错误日志输出
		&utils.LoggerWriter{Logger: middleware.GetLogger(), Level: "error"}, // 同步到 zap error
	)

	// 加载配置文件
	err = config.LoadConfig()
	if err != nil {
		middleware.GetLogger().Error("配置文件加载错误: %v", zap.Error(err))
		return
	}
	// 初始化所有模块
	err = InitializeAll()
	if err != nil {
		middleware.GetLogger().Error("初始化模块错误: %v", zap.Error(err))
		return
	}

	// 初始化路由
	r := gin.Default()
	v1.SetupRoutes(r, Engine)

	// 启动服务
	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		middleware.GetLogger().Error("服务启动错误: %v", zap.Error(err))
		return
	}
}
