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
	order.Find()
	//if err != nil {
	//	return 0 ,fmt.Errorf("查询订单失败:%v",err)
	//}
	part := model.Partlist{
		Hot: hot,
	}
	list,_ := part.Find(10,true)
	//if err != nil {
	//	return 0 ,fmt.Errorf("查询热门列表失败:%v",err)
	//}
	fmt.Print("%+v",list)
	return list,nil
}