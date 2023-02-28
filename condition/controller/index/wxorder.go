package index

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/logic/user"
	"gopartsrv/condition/model"
	"gopartsrv/public/consts"
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

//Order 下单
func Order(c *gin.Context) {
	openid := c.Query("openid")
	userid := c.Query("userid")
	amount := c.Query("amount")
	amounts, err := strconv.ParseInt(amount, 10, 64)
	res,err := wxpay.CreatOrder(openid,amounts)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err,"code":2001, "msg": "请求失败", "data": ""})
		return
	}
	fmt.Println(res)
	fmt.Println(err)
	//wxpay.CallBack()
	list, err := user.UserInfo(userid,openid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err,"code":2002, "msg": "请求失败", "data": ""})
		return
	}
	fmt.Println(list)
	order:=model.Order{
		Id: list.Id,
		Openid: list.Openid,
		Amount: amounts,
		Createtime: time.Now().Format(consts.FORMATDATELONG),
		Updatetime: time.Now().Format(consts.FORMATDATELONG),
	}
	err = order.Create()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err,"code":2003, "msg": "请求失败", "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "请求成功", "data": res})
	return
}
func OrderBack(c *gin.Context)  {
	handler := wxpay.CallBack()
	var content interface{}
	handler.ParseNotifyRequest(context.Background(),c.Request,&content)
	fmt.Println(content)
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data": 1})
	return
}

