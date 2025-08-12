package middleware

import (
	"fmt"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	// 创建按日期轮转的日志写入器
	writer, err := rotatelogs.New(
		"logs/cmorder2db-%Y-%m-%d.log",            // 按日期命名
		rotatelogs.WithMaxAge(365*24*time.Hour),   // 保留365天
		rotatelogs.WithRotationTime(24*time.Hour), // 每24小时轮转
	)

	if err != nil {
		fmt.Printf("logger异常 %s", err.Error())
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 级别大写

	// 创建文件输出核心
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(writer),
		zapcore.InfoLevel,
	)

	// 创建日志
	Logger = zap.New(core, zap.AddCaller())
}

func GetLogger() *zap.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}
