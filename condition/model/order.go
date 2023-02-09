package model

import (
	"gopartsrv/utils/db"
)

//获取连接
func init() {
	dbs = db.DBTeamMap["mini"] //同一个db
}

//订单
type Order struct {
	Id         string `json:"id"` //banner的id
	Userid   int `json:"userid"`
	Openid string `json:"openid"`
	Amount float64 `json:"amount"`
	Createtime   string `json:"createtime"` //创建时间
	Updatetime string `json:"updatetime"`
	Deletetime string `json:"deletetime"`
}
//获取表名
func (u *Order) TableName() string {
	return "order"
}

//查询内容(id 查询)
func (u *Order) Find() (*[]Order, error) {
	list := &[]Order{}
	sqls := dbs.Table(u.TableName())
	if u.Id != "" {
		sqls = sqls.Where("id = ?", u.Id)
	}
	if u.Openid != "" {
		sqls = sqls.Where("openid = ?", u.Openid)
	}
	err := sqls.Find(list).Error
	if err != nil {
		return &[]Order{}, err
	}
	return list, nil
}
//查询开始和结束时间
func FindTime()  {

}

