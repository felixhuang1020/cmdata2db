package controller

import (
	"cmdata2db/internal/model"
	"cmdata2db/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(UserService *service.UserService) *UserController {
	return &UserController{UserService: UserService}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

type CreateUserRequest struct {
	UserID      int64  `json:"userID" binding:"required"`                // 对应模型的UserID，必填
	Password    string `json:"password" binding:"required,min=6,max=50"` // 密码，必填，长度6-50（匹配模型size:50）
	UserName    string `json:"userName" binding:"max=30"`                // 用户名，最长30（匹配模型size:30）
	Email       string `json:"email" binding:"omitempty,email,max=50"`   // 邮箱，可选，符合格式，最长50
	PhoneNumber string `json:"phoneNumber" binding:"omitempty,max=20"`   // 手机号，可选，最长20
	Sex         string `json:"sex" binding:"omitempty,oneof=男 女 未知"`     // 性别，可选，限制可选值
	Remark      string `json:"remark" binding:"omitempty,max=500"`       // 备注，可选，最长500
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 2. 转换请求参数为User模型
	user := &model.User{
		Id:          req.UserID,
		Password:    req.Password, // 注意：实际项目中这里需要加密处理（如bcrypt）
		UserName:    req.UserName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Sex:         req.Sex,
		Remark:      req.Remark,
	}

	if err := uc.UserService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建用户失败: " + err.Error(),
		})
		return
	}

	// 4. 返回成功响应（只返回必要字段，避免敏感信息泄露）
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建用户成功",
		"data": gin.H{
			"id":       user.Id,       // 模型中的主键Id
			"userID":   user.UserID,   // 业务用户ID
			"userName": user.UserName, // 用户名
			"email":    user.Email,
		},
	})
}

func (uc *UserController) GetProducts(c *gin.Context) {
	products, err := uc.UserService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch productss"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

type CreateProductRequest struct {
	ProductID     string  `json:"productID" binding:"required"`                // 对应模型的ProductID，必填
	ProductName   string  `json:"productName" binding:"required,min=6,max=50"` // 产品名称，必填，长度6-50（匹配模型size:50）
	Description   string  `json:"description" binding:"max=500"`
	Category      string  `json:"category" binding:"required"`
	SubCategory   string  `json:"subCategory" binding:"omitempty"`
	Brand         string  `json:"brand" binding:"required,min=2,max=20"`
	Price         float64 `json:"price" binding:"required,min=0"`
	OriginalPrice float64 `json:"originalPrice" binding:"omitempty,min=0"`
	Currency      string  `json:"currency" binding:"omitempty,oneof=CNY USD"`
	Stock         int     `json:"stock" binding:"required,min=0"`
	IsInStock     bool    `json:"isInStock" binding:"required"`
}

func (uc *UserController) CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	// 2. 转换请求参数为Product模型
	product := &model.Product{
		ID:            req.ProductID,
		Name:          req.ProductName,
		Description:   req.Description,
		Category:      req.Category,
		SubCategory:   req.SubCategory,
		Brand:         req.Brand,
		Price:         req.Price,
		OriginalPrice: req.OriginalPrice,
		Currency:      req.Currency,
		Stock:         req.Stock,
		IsInStock:     req.IsInStock,
	}
	if err := uc.UserService.CreateProducts(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建产品失败: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建产品成功",
		"data": gin.H{
			"id":            product.ID,   // 模型中的主键Id
			"productID":     product.ID,   // 业务产品ID
			"productName":   product.Name, // 产品名称
			"description":   product.Description,
			"category":      product.Category,
			"subCategory":   product.SubCategory,
			"brand":         product.Brand,
			"price":         product.Price,
			"originalPrice": product.OriginalPrice,
			"currency":      product.Currency,
			"stock":         product.Stock,
		},
	})
}

func (uc *UserController) PurchaseProduct(c *gin.Context) {
	productID := c.Query("productID")
	number, err := strconv.Atoi(c.Query("number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: number 必须是一个整数" + err.Error(),
		})
		return
	}
	if err := uc.UserService.PurchaseProduct(productID, number); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "购买产品失败: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "购买产品成功",
	})
}

func (uc *UserController) DeleteProduct(c *gin.Context) {
	productID := c.Query("productID")
	if err := uc.UserService.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "删除产品失败: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除产品成功",
	})
}
