package index

import (
	"fmt"
	"gopartsrv/condition/model"
	"gopartsrv/public/consts"
	"time"
)

//Hotlist 热门列表
func Hotlist(userId ,openid string,hot int) (*[]model.Partlist,error) {
	user:=model.Users{
		Id: userId,
		Openid: openid,
	}
	users,err:=user.Find()
	if err != nil {
		return &[]model.Partlist{} ,fmt.Errorf("查询用户失败:%v",err)
	}
	part := model.Partlist{
		Hot: hot,
	}
	buy := false
	local, _ := time.LoadLocation("Local")
	locationDatetime, _ := time.ParseInLocation(consts.FORMATDATELONG, users.Cuttime, local)
	if locationDatetime.Unix() >= time.Now().Unix() {
		buy = true
	}
	list,_ := part.Find(5,buy)
	for i, item := range *list {
		item.Buy = buy
		item.Img = users.Address
		(*list)[i] = item
	}
	if err != nil {
		return list ,fmt.Errorf("查询热门列表失败:%v",err)
	}
	return list,nil
}

func Partlist(userId,openid,search,city,area string,page,pageSize int)(*[]model.Partlist,error)  {
	user:=model.Users{
		Id: userId,
		Openid: openid,
	}
	users,err:=user.Find()
	if err != nil {
		return &[]model.Partlist{} ,fmt.Errorf("查询用户失败:%v",err)
	}
	part := model.Partlist{
		City: city,
		Area: area,
		Content: search,
		Status: 1,
	}
	buy := false
	local, _ := time.LoadLocation("Local")
	locationDatetime, _ := time.ParseInLocation(consts.FORMATDATELONG, users.Cuttime, local)
	if locationDatetime.Unix() >= time.Now().Unix() {
		buy = true
	}
	list,_ := part.Find2Search(page,pageSize,buy)
	for i, item := range *list {
		item.Buy = buy
		item.Img = users.Address
		(*list)[i] = item
	}
	if err != nil {
		return list ,fmt.Errorf("查询热门列表失败:%v",err)
	}
	return list,nil
}