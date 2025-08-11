package service

import (
	"cmdata2db/internal/model"
	"cmdata2db/internal/repository"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(engine *gorm.DB) *UserService {
	return &UserService{userRepo: repository.NewUserRepository(engine)}
}

func (us *UserService) GetUsers() ([]*model.User, error) {
	return us.userRepo.GetUsers()
}

func (us *UserService) CreateUser(user *model.User) error {
	return us.userRepo.CreateUser(user)
}

func (us *UserService) GetProducts() ([]*model.Product, error) {
	return us.userRepo.GetProducts()
}

func (us *UserService) CreateProducts(product *model.Product) error {
	return us.userRepo.CreateProduct(product)
}

func (us *UserService) PurchaseProduct(productID string, number int) error {
	return us.userRepo.PurchaseProduct(productID, number)
}

func (us *UserService) DeleteProduct(productID string) error {
	return us.userRepo.DeleteProduct(productID)
}
