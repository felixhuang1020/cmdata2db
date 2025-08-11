package app

import (
	"cmdata2db/config"
	v1 "cmdata2db/internal/api/v1"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Start() {
	// 加载配置文件
	err := config.LoadConfig()
	if err != nil {
		log.Error("配置文件加载错误: %v", err)
		return
	}

	// 初始化所有模块
	err = InitializeAll()
	if err != nil {
		log.Error("模块初始化错误: %v", err)
		return
	}

	// 初始化路由
	r := gin.Default()
	v1.SetupRoutes(r, Engine)

	// 启动服务
	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		log.Error("服务启动错误: %v", err)
		return
	}
}
