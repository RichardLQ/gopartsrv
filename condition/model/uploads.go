package model

import (
	"gopartsrv/utils/db"
)

//获取连接
func init() {
	dbs = db.DBTeamMap["work"] //同一个db
}

//上传信息
type Uploads struct {
	Id         string `json:"id"`
	Openid     string `json:"openid"`
	Address    string `json:"address"`
	Type       int    `json:"type"`
	Createtime string `json:"createtime"`
}

type UploadFile struct {
	Id         string `json:"id"`
	Openid     string `json:"openid"`
	Address    string `json:"address"`
	Type       int    `json:"type"`
	Createtime string `json:"createtime"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}

//获取表名
func (u *UploadFile) TableName() string {
	return "uploads"
}

//查询内容
func (u *UploadFile) FindArray() (*[]Uploads, error) {
	list := &[]Uploads{}
	sqls := dbs.Table(u.TableName())
	sqls = sqls.Where("openid = ?", u.Openid)
	sqls = sqls.Where("type = ?", u.Type).Order("createtime desc").Limit(u.PageSize).Offset((u.Page - 1) * u.PageSize)
	err := sqls.Find(list).Error
	if err != nil {
		return &[]Uploads{}, err
	}
	return list, nil
}

//查询内容
func (u *UploadFile) Find() (*[]Uploads, error) {
	list := &[]Uploads{}
	sqls := dbs.Table(u.TableName())
	if u.Id != "" {
		sqls = sqls.Where("id = ?", u.Id)
	}
	if u.Openid != "" {
		sqls = sqls.Where("openid = ?", u.Openid)
	}
	err := sqls.Find(list).Error
	if err != nil {
		return &[]Uploads{}, err
	}
	return list, nil
}

//添加数据
func (u *UploadFile) Create() error {
	err := dbs.Table(u.TableName()).
		Create(u).Error
	if err != nil {
		return err
	}
	return err
}

//删除数据
func (u *UploadFile) Delete() (int, error) {
	err := dbs.Table(u.TableName()).Where("id = ?", u.Id).Delete(&Users{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}
