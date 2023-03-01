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
	Id         int64  `json:"id"` //banner的id
	Userid     int    `json:"userid"`
	Openid     string `json:"openid"`
	Amount     int64  `json:"amount"`
	Tradeno string `json:"tradeno"` //订单id
	Transactionid string `json:"transactionid"`
	Status     int    `json:"status"`
	Createtime string `json:"createtime"` //创建时间
	Updatetime string `json:"updatetime"`
	Deletetime string `json:"deletetime"`
}

//获取表名
func (u *Order) TableName() string {
	return "order"
}

//Find 查询内容(id 查询)
func (u *Order) Find() (*[]Order, error) {
	list := &[]Order{}
	sqls := dbs.Table(u.TableName())
	if u.Id != 0 {
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

//Find 查询内容(id 查询)
func (u *Order) FindOne() (*Order, error) {
	list := &Order{}
	sqls := dbs.Table(u.TableName())
	if u.Id != 0 {
		sqls = sqls.Where("id = ?", u.Id)
	}
	if u.Openid != "" {
		sqls = sqls.Where("openid = ?", u.Openid)
	}
	err := sqls.First(list).Error
	if err != nil {
		return &Order{}, err
	}
	return list, nil
}

func (u *Order) Create() (id int64, err error) {
	sql := dbs.Table(u.TableName())
	id = u.Id
	if u.Id != 0 {
		err = sql.Save(u).Error
		return
	}
	result := sql.Create(u).RowsAffected
	return result, nil
}

func (u *Order) Update(types string) error {
	err := dbs.Table(u.TableName()).Where("tradeno = ?",u.Tradeno).
		Updates(Order{
			Openid: u.Openid,
			Transactionid:u.Transactionid,
			Updatetime: u.Updatetime,
	}).Error
	if err != nil {
		return err
	}
	if u.Status == 2 {
		list,err := u.FindOne()
		if err!= nil {
			user:=&Users{
				Openid: list.Openid,
			}
			user.UpdateVip(types)
		}
	}
	return nil
}


//查询开始和结束时间
func FindTime() {

}
