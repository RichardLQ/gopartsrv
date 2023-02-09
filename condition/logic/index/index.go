package index

import (
	"fmt"
	"gopartsrv/condition/model"
)

//Hotlist 热门列表
func Hotlist(userId,openid string,hot int) (*[]model.Partlist,error) {
	order:=model.Order{
		Id: userId,
		Openid: openid,
	}
	orders,err:=order.Find()
	if err != nil {
		return &[]model.Partlist{} ,fmt.Errorf("查询订单失败:%v",err)
	}
	part := model.Partlist{
		Hot: hot,
	}
	buy := false
	if len(*orders)>0 {
		buy = true
	}
	list,_ := part.Find(10,buy)
	if err != nil {
		return list ,fmt.Errorf("查询热门列表失败:%v",err)
	}
	return list,nil
}

func Partlist()  {
	
}