package index

import (
	"context"
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/model"
	"gopartsrv/public/consts"
	"gopartsrv/utils/mini"
	"gopartsrv/utils/wxpay"
	"net/http"
	"strconv"
	"time"
)

type OrderBody struct {
	Mchid       string  `json:"mchid"`
	OutTradeNo  string  `json:"out_trade_no"`
	Appid       string  `json:"appid"`
	Description string  `json:"description"`
	NotifyUrl   string  `json:"notify_url"`
	Amount      Amount  `json:"amount"`
	Payer       PayUser `json:"payer"`
}

type Amount struct {
	Total    int    `json:"total"`
	Currency string `json:"currency"`
}
type PayUser struct {
	Openid string `json:"openid"`
}

//CreatOrder 下单
func CreatOrder(c *gin.Context) {
	openid := c.Query("openid")
	userid := c.Query("userid")
	userids, err := strconv.Atoi(userid)
	amount := c.Query("amount")
	types := c.Query("type")
	appid := mini.APPID
	if types == "2" {
		appid = mini.LMP_APPID
	}
	amounts, err := strconv.ParseInt(amount, 10, 64)
	res,rid, err := wxpay.CreatOrder(openid,appid, amounts,userids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "code": 2001, "msg": "请求失败","rid":0, "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "请求成功", "rid":rid,"data": res})
	return
}

//PayStatus 支付
func PayStatus(c *gin.Context)  {
	rid := c.Query("rid")
	types := c.Query("type")//type:0：月，1：季度，2：年
	if rid == ""{
		c.JSON(http.StatusOK, gin.H{"errs": "", "code": 2003, "msg": "请求订单找不到", "data": ""})
		return
	}
	status := c.Query("status")
	statuss, _ := strconv.Atoi(status)
	rids, _ := strconv.ParseInt(rid, 10, 64)
	order := model.Order{
		Id:         rids,
		Status: statuss,
		Createtime: time.Now().Format(consts.FORMATDATELONG),
	}
	err:=order.Update(types)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "code": 2003, "msg": "请求失败", "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "请求成功", "data": ""})
}

func OrderBack(c *gin.Context) {
	handler := wxpay.CallBack()
	var content interface{}
	handler.ParseNotifyRequest(context.Background(), c.Request, &content)
	openid := content.(map[string]interface{})["payer"].(map[string]interface{})["openid"].(string)
	total := content.(map[string]interface{})["amount"].(map[string]interface{})["total"].(float64)
	out_trade_no := content.(map[string]interface{})["out_trade_no"].(string)
	transaction_id := content.(map[string]interface{})["transaction_id"].(string)
	success_time := content.(map[string]interface{})["success_time"].(string)
	status := content.(map[string]interface{})["trade_state"].(string)
	ts, _ := time.Parse(time.RFC3339, success_time)
	types := "0"
	if 1000< total && total<= 10000 {
		types = "1"
	}
	if 10000< total && total<= 100000 {
		types = "2"
	}
	//openid := "oBzet53gPZSisPu4XgCWNCn8pm68"
	//out_trade_no := "d3500d4f-7bfb-499b"
	//transaction_id := "4200001756202303021462235849"
	//status:= "SUCCESS"
	//types := "0"
	order:= model.Order{
		Openid: openid,
		Tradeno: out_trade_no,
		Status: 1,
		Transactionid: transaction_id,
		//Updatetime: "2023-03-02 12:10:41",
		Updatetime: ts.Format(consts.FORMATDATELONG),
	}
	if status == "SUCCESS" {
		order.Status = 2
	}
	order.Update(types)
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data": 1})
	return
}
