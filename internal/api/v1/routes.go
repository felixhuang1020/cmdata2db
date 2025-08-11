package v1

import (
	"cmdata2db/internal/controller"
	"cmdata2db/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, engine *gorm.DB) {
	// 定义用户路由组
	user := r.Group("/user")
	{
		// 创建 UserService 实例
		UserService := service.NewUserService(engine)
		// 创建 UserController 实例
		UserController := controller.NewUserController(UserService)

		user.GET("/", UserController.GetUsers)
		user.POST("/", UserController.CreateUser)
	}

	product := r.Group("/product")
	{
		UserService := service.NewUserService(engine)
		UserController := controller.NewUserController(UserService)
		product.GET("/", UserController.GetProducts)
		product.POST("/", UserController.CreateProduct)
	}

	purchase := r.Group("/purchase")
	{
		UserService := service.NewUserService(engine)
		UserController := controller.NewUserController(UserService)
		purchase.POST("/", UserController.PurchaseProduct)
	}

	deleteproduct := r.Group("/deleteproduct")
	{
		UserService := service.NewUserService(engine)
		UserController := controller.NewUserController(UserService)
		deleteproduct.POST("/", UserController.DeleteProduct)
	}

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
