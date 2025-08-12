package v1

import (
	"cmdata2db/internal/controller"
	"cmdata2db/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, engine *gorm.DB) {
	// 定义用户路由组
	order := r.Group("/order")
	{
		OrderService := service.NewOrderService(engine)
		OrderController := controller.NewOrderController(OrderService)
		order.POST("/", OrderController.GetOrders)
	}

	batchorder := r.Group("/batchorder")
	{
		OrderService := service.NewOrderService(engine)
		OrderController := controller.NewOrderController(OrderService)
		batchorder.POST("/", OrderController.SaveBatchOrderData)
	}
}
