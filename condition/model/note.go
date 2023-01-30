package model

import (
	"gopartsrv/utils/db"
)

//获取连接
func init() {
	dbs = db.DBTeamMap["work"] //同一个db
}

//用户信息
type Note struct {
	Id         int64 `json:"id"`
	Openid     string `json:"openid"`
	Content    string `json:"content"`
	Updatetime string `json:"updatetime"`
	Createtime string `json:"createtime"`
}

//获取表名
func (u *Note) TableName() string {
	return "note"
}

//查询内容
func (u *Note) Find() (*Note, error) {
	list := &Note{}
	sqls := dbs.Table(u.TableName())
	if u.Id != 0 {
		sqls = sqls.Where("id = ?", u.Id)
	}
	if u.Openid != "" && u.Createtime != ""{
		sqls = sqls.Where("openid = ?", u.Openid).Where("createtime = ?", u.Createtime)
	}
	err := sqls.Find(list).Error
	if err != nil {
		return &Note{}, err
	}
	return list, nil
}

func (u *Note) FindArray(page,pageSize int64) (*[]Note, error) {
	list := &[]Note{}
	sqls := dbs.Table(u.TableName())
	sqls = sqls.Where("openid = ?", u.Openid).Order("createtime desc").Limit(pageSize).Offset((page - 1) * pageSize)
	err := sqls.Find(list).Error
	if err != nil {
		return &[]Note{}, err
	}
	return list, nil
}

//添加数据
func (u *Note) Create() (Note,error) {
	err := dbs.Table(u.TableName()).
		Create(u).Error
	if err != nil {
		return Note{},err
	}
	return *u,err
}

//删除数据
func (u *Note) Delete() (int, error) {
	err := dbs.Table(u.TableName()).Where("id = ?", u.Id).Delete(&Note{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (u *Note) Update() (Note, error){
	err := dbs.Table(u.TableName()).Where("id = ?", u.Id).Updates(u).Error
	if err != nil {
		return Note{}, err
	}
	return *u, nil
}

