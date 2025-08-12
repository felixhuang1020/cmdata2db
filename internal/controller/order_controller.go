package controller

import (
	"cmdata2db/config"
	"cmdata2db/internal/middleware"
	"cmdata2db/internal/model"
	"cmdata2db/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type OrderController struct {
	OrderService *service.OrderService
}

func NewOrderController(OrderService *service.OrderService) *OrderController {
	return &OrderController{OrderService: OrderService}
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	orders, err := oc.OrderService.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		middleware.GetLogger().Error("Failed to fetch orders", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

type TbCustInstanceProcessRequest struct {
	OrderID             string `json:"orderid,omitempty"`
	TaskID              string `json:"taskid,omitempty"`
	Creatime            string `json:"creatime,omitempty"`
	Dealtime            string `json:"dealtime,omitempty"`
	Answertime          string `json:"answertime,omitempty"`
	Dealduration        string `json:"dealduration,omitempty"`
	Answerduration      string `json:"answerduration,omitempty"`
	Creatuser           string `json:"creatuser,omitempty"`
	Dealuser            string `json:"dealuser,omitempty"`
	Dealcontent         string `json:"dealcontent,omitempty"`
	Dealprocess         string `json:"dealprocess,omitempty"`
	Operate             string `json:"operate,omitempty"`
	State               string `json:"state,omitempty"`
	Dealsla             string `json:"dealsla,omitempty"`
	Answersla           string `json:"answersla,omitempty"`
	Dealdeadline        string `json:"dealdeadline,omitempty"`
	Answerdeadline      string `json:"answerdeadline,omitempty"`
	Dealtimeout         string `json:"dealtimeout,omitempty"`
	Answertimeout       string `json:"answertimeout,omitempty"`
	Pendduration        string `json:"pendduration,omitempty"`
	Dealtimeoutcause    string `json:"dealtimeoutcause,omitempty"`
	Responsetimeoutcase string `json:"responsetimeoutcase,omitempty"`
	Transferreason      string `json:"transferreason,omitempty"`
}

type APIResponse struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    []TbCustInstanceProcessRequest `json:"data"`
}

func (oc *OrderController) SaveBatchOrderData(c *gin.Context) {
	start := time.Now()
	middleware.GetLogger().Info("开始获取订单数据")
	response, err := http.Post(config.Conf.App.RequestUrl, "application/json", nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		middleware.GetLogger().Error("获取订单数据失败: ", zap.Error(err))
		return
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			middleware.GetLogger().Warn("关闭HTTP响应体失败: ", zap.Error(err))
		}
	}()

	if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "外部API返回错误状态: " + response.Status,
		})
		middleware.GetLogger().Error("外部API返回错误状态: ", zap.Int("statusCode", response.StatusCode))
		return
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "解析API响应失败",
			"body": apiResponse,
		})
		middleware.GetLogger().Error("解析API响应失败: ", zap.Error(err))
		return
	}

	// 处理多个订单数据
	var orders []*model.Tb_cust_instanceprocess
	for _, item := range apiResponse.Data {
		order := &model.Tb_cust_instanceprocess{
			OrderID:             item.OrderID,
			TaskID:              item.TaskID,
			Creatime:            item.Creatime,
			Dealtime:            item.Dealtime,
			Answertime:          item.Answertime,
			Dealduration:        item.Dealduration,
			Answerduration:      item.Answerduration,
			Creatuser:           item.Creatuser,
			Dealuser:            item.Dealuser,
			Dealcontent:         item.Dealcontent,
			Dealprocess:         item.Dealprocess,
			Operate:             item.Operate,
			State:               item.State,
			Dealsla:             item.Dealsla,
			Answersla:           item.Answersla,
			Dealdeadline:        item.Dealdeadline,
			Answerdeadline:      item.Answerdeadline,
			Dealtimeout:         item.Dealtimeout,
			Answertimeout:       item.Answertimeout,
			Pendduration:        item.Pendduration,
			Dealtimeoutcause:    item.Dealtimeoutcause,
			Responsetimeoutcase: item.Responsetimeoutcase,
			Transferreason:      item.Transferreason,
			WriteTime:           time.Now(),
		}
		orders = append(orders, order)
	}

	// 检查是否有数据
	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "没有数据需要保存",
			"count": 0,
		})
		return
	}
	middleware.GetLogger().Info("开始保存订单数据", zap.Int("数量", len(orders)))

	// 保存所有订单数据
	batches := lo.Chunk(orders, config.Conf.App.Batch)
	for i, batch := range batches {
		middleware.GetLogger().Info(fmt.Sprintf("处理第 %d 批，共 %d 条数据", i+1, len(batch)))
		if err := oc.OrderService.SaveBatchOrderData(batch); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "保存数据失败: ",
			})
			middleware.GetLogger().Error("保存数据失败: ", zap.Error(err))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "数据保存成功",
			"count": len(batch),
			"data":  batch,
		})
	}

	middleware.GetLogger().Info("订单数据保存完成", zap.Duration("耗时", time.Since(start)))
}
