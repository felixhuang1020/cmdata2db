package model

import (
	"time"
)

type Tb_cust_instanceprocess struct {
	OrderID             string    `gorm:"column:orderid"`
	TaskID              string    `gorm:"column:taskid"`
	Creatime            string    `gorm:"column:creatime"`
	Dealtime            string    `gorm:"column:dealtime"`
	Answertime          string    `gorm:"column:answertime"`
	Dealduration        string    `gorm:"column:dealduration"`   // 使用指针类型
	Answerduration      string    `gorm:"column:answerduration"` // 使用指针类型
	Creatuser           string    `gorm:"column:creatuser"`
	Dealuser            string    `gorm:"column:dealuser"`
	Dealcontent         string    `gorm:"column:dealcontent"`
	Dealprocess         string    `gorm:"column:dealprocess"`
	Operate             string    `gorm:"column:operate"`
	State               string    `gorm:"column:state"`
	Dealsla             string    `gorm:"column:dealsla"`
	Answersla           string    `gorm:"column:answersla"`
	Dealdeadline        string    `gorm:"column:dealdeadline"`
	Answerdeadline      string    `gorm:"column:answerdeadline"`
	Dealtimeout         string    `gorm:"column:dealtimeout"`   // 使用指针类型
	Answertimeout       string    `gorm:"column:answertimeout"` // 使用指针类型
	Pendduration        string    `gorm:"column:pendduration"`  // 使用指针类型
	Dealtimeoutcause    string    `gorm:"column:dealtimeoutcause"`
	Responsetimeoutcase string    `gorm:"column:responsetimeoutcase"`
	Transferreason      string    `gorm:"column:transferreason"`
	WriteTime           time.Time `gorm:"column:write_time"`
}
