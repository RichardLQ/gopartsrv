package model

import (
	"github.com/jinzhu/gorm"
	"gopartsrv/utils/db"
)

var (
	dbs *gorm.DB
)

//获取连接
func init() {
	dbs = db.DBTeamMap["mini"] //同一个db
}

//用户信息
type Users struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Openid     string `json:"openid"`
	Gender     string `json:"gender"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	Nickname   string `json:"nickname"`
	Desc       string `json:"desc"`
	Updatetime string `json:"updatetime"`
	Createtime string `json:"createtime"`
}

//获取表名
func (u *Users) TableName() string {
	return "users"
}

//查询内容
func (u *Users) Find() (*Users, error) {
	list := &Users{}
	sqls := dbs.Table(u.TableName())
	if u.Id != "" {
		sqls = sqls.Where("id = ?", u.Id)
	}
	if u.Openid != "" {
		sqls = sqls.Where("openid = ?", u.Openid)
	}
	err := sqls.Find(list).Error
	if err != nil {
		return &Users{}, err
	}
	return list, nil
}

//添加数据
func (u *Users) Create() error {
	err := dbs.Table(u.TableName()).
		Create(u).Error
	if err != nil {
		return err
	}
	return err
}

//删除数据
func (u *Users) Delete() (int, error) {
	err := dbs.Table(u.TableName()).Where("id = ?", u.Id).Delete(&Users{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (u *Users) Updates() (int, error) {
	sqls := dbs.Table(u.TableName())
	if u.Id != "" {
		sqls = sqls.Where("id = ?", u.Id)
	}
	if u.Openid != "" {
		sqls = sqls.Where("openid = ?", u.Openid)
	}
	err := sqls.Updates(u).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}
