package service

import (
	"cmdata2db/internal/model"
	"cmdata2db/internal/repository"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
}

func NewOrderService(engine *gorm.DB) *OrderService {
	return &OrderService{orderRepo: repository.NewOrderRepository(engine)}
}

func (os *OrderService) GetOrders() ([]*model.Tb_cust_instanceprocess, error) {
	return os.orderRepo.GetOrders()
}

func (os *OrderService) SaveBatchOrderData(orders []*model.Tb_cust_instanceprocess) error {
	return os.orderRepo.SaveBatchOrderData(orders)
}
