package main

import (
	"cmdata2db/internal/app"
	"cmdata2db/internal/middleware"
)

func main() {
	middleware.GetLogger().Info("服务启动成功")
	app.Start()
}
