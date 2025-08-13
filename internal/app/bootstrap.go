package app

import (
	"cmdata2db/config"
	v1 "cmdata2db/internal/api/v1"
	"cmdata2db/internal/controller"
	"cmdata2db/internal/middleware"
	"cmdata2db/internal/service"
	"cmdata2db/internal/utils"
	"fmt"
	"io"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
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

	// 初始化定时任务
	cronSpec := config.Conf.App.CronSpec
	if cronSpec == "" {
		middleware.GetLogger().Warn("未配置cron表达式，定时任务未启动")
	} else {
		c := cron.New()
		_, err := c.AddFunc(cronSpec, func() {
			middleware.GetLogger().Info("定时任务开始执行")
			// 创建模拟上下文
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			// 实例化控制器
			orderService := service.NewOrderService(Engine)
			orderController := controller.NewOrderController(orderService)
			// 执行任务
			orderController.SaveBatchOrderData(ctx)
		})
		if err != nil {
			middleware.GetLogger().Error("定时任务添加失败", zap.Error(err))
		} else {
			c.Start()
			defer c.Stop()
			middleware.GetLogger().Info("定时任务已启动", zap.String("cron_spec", cronSpec))
		}
	}

	// 启动服务
	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		middleware.GetLogger().Error("服务启动错误: %v", zap.Error(err))
		return
	}
}
