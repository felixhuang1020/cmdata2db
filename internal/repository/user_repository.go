package repository

import (
	"cmdata2db/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	engine *gorm.DB
}

func NewUserRepository(engine *gorm.DB) *UserRepository {
	return &UserRepository{engine: engine}
}

// GetUsers 获取所有用户
func (r *UserRepository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := r.engine.Find(&users).Error
	return users, err
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.engine.Create(user).Error
}

func (r *UserRepository) GetProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := r.engine.Find(&products).Error
	return products, err
}

func (r *UserRepository) CreateProduct(product *model.Product) error {
	return r.engine.Create(product).Error
}

func (r *UserRepository) PurchaseProduct(productID string, number int) error {
	return r.engine.Model(&model.Product{}).Where("id = ? AND stock > 0", productID).Update("stock", gorm.Expr("stock - ?", number)).Error
}

func (r *UserRepository) DeleteProduct(productID string) error {
	return r.engine.Model(&model.Product{}).Where("id = ?", productID).Delete(&model.Product{}).Error
}
