package controller

import (
	"cmdata2db/config"
	"cmdata2db/internal/model"
	"cmdata2db/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
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
	response, err := http.Post("http://127.0.0.1:4523/m2/6021227-5710205-default/286511927", "application/json", nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Warnf("关闭HTTP响应体失败: %v", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "外部API返回错误状态: " + response.Status,
		})
		return
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "解析API响应失败: " + err.Error(),
			"body": apiResponse,
		})
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

	// 保存所有订单数据
	batches := lo.Chunk(orders, config.Conf.App.Batch)
	for i, batch := range batches {
		fmt.Printf("处理第 %d 批，共 %d 条数据\n", i+1, len(batch))
		if err := oc.OrderService.SaveBatchOrderData(batch); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "保存数据失败: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "数据保存成功",
			"count": len(batch),
			"data":  batch,
		})
	}

}
