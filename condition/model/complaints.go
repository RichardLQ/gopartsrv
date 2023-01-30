package model

import "gopartsrv/utils/db"

//获取连接
func init() {
	dbs = db.DBTeamMap["work"] //同一个db
}

//投诉信息
type Complaints struct {
	Id         string `json:"id"`
	Openid     string `json:"openid"`
	Content       string `json:"content"`
	Createtime string `json:"createtime"`
}

//获取表名
func (u *Complaints) TableName() string {
	return "complaints"
}

//添加数据
func (u *Complaints) Create() error {
	err := dbs.Table(u.TableName()).
		Create(u).Error
	if err != nil {
		return err
	}
	return err
}

