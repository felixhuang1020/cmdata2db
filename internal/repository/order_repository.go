package repository

import (
	"cmdata2db/internal/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	engine *gorm.DB
}

func NewOrderRepository(engine *gorm.DB) *OrderRepository {
	return &OrderRepository{engine: engine}
}

func (r *OrderRepository) GetOrders() ([]*model.Tb_cust_instanceprocess, error) {
	var orders []*model.Tb_cust_instanceprocess
	err := r.engine.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) SaveBatchOrderData(orders []*model.Tb_cust_instanceprocess) error {
	return r.engine.Create(orders).Error
}
