package controller

import (
	"cmdata2db/internal/model"
	"cmdata2db/internal/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "外部API返回错误状态: " + response.Status,
		})
		return
	}

	var apiData APIResponse
	if err := json.NewDecoder(response.Body).Decode(&apiData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "解析API响应失败: " + err.Error(),
			"body": apiData, // 显示实际响应内容
		})
		return
	}

	orders := []*model.Tb_cust_instanceprocess{
		{
			OrderID:             apiData.Data[0].OrderID,
			TaskID:              apiData.Data[0].TaskID,
			Creatime:            apiData.Data[0].Creatime,
			Dealtime:            apiData.Data[0].Dealtime,
			Answertime:          apiData.Data[0].Answertime,
			Dealduration:        apiData.Data[0].Dealduration,
			Answerduration:      apiData.Data[0].Answerduration,
			Creatuser:           apiData.Data[0].Creatuser,
			Dealuser:            apiData.Data[0].Dealuser,
			Dealcontent:         apiData.Data[0].Dealcontent,
			Dealprocess:         apiData.Data[0].Dealprocess,
			Operate:             apiData.Data[0].Operate,
			State:               apiData.Data[0].State,
			Dealsla:             apiData.Data[0].Dealsla,
			Answersla:           apiData.Data[0].Answersla,
			Dealdeadline:        apiData.Data[0].Dealdeadline,
			Answerdeadline:      apiData.Data[0].Answerdeadline,
			Dealtimeout:         apiData.Data[0].Dealtimeout,
			Answertimeout:       apiData.Data[0].Answertimeout,
			Pendduration:        apiData.Data[0].Pendduration,
			Dealtimeoutcause:    apiData.Data[0].Dealtimeoutcause,
			Responsetimeoutcase: apiData.Data[0].Responsetimeoutcase,
			Transferreason:      apiData.Data[0].Transferreason,
			WriteTime:           time.Now(),
		},
	}

	if err := oc.OrderService.SaveBatchOrderData(orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "保存数据失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据保存成功",
		"data": orders,
	})

}
