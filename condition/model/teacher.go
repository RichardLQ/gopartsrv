package model

import (
	"gopartsrv/utils/db"
)

//获取连接
func init() {
	dbs = db.DBTeamMap["work"] //同一个db
}

//老师信息
type Teacher struct {
	Id         string `json:"id"`         //老师id
	Nickname   string `json:"nickname"`   //昵称
	Name       string `json:"name"`       //名字
	Address    string `json:"address"`    //头像地址
	Desc       string `json:"desc"`       //简介
	Phone      int64  `json:"phone"`      //电话号码
	Qrcode     string `json:"qrcode"`     //二维码
	Type       int64  `json:"type"`       //舞蹈类型
	Updateime  int64  `json:"updateime"`  //更新时间
	Createtime int64  `json:"createtime"` //创建时间
}

//老师请求体
type TeacherType struct {
	Id         string `json:"id"`         //老师id
	Nickname   string `json:"nickname"`   //昵称
	Name       string `json:"name"`       //名字
	Address    string `json:"address"`    //头像地址
	Desc       string `json:"desc"`       //简介
	Phone      int64  `json:"phone"`      //电话号码
	Qrcode     string `json:"qrcode"`     //二维码
	Type       int64  `json:"type"`       //舞蹈类型
	Updateime  int64  `json:"updateime"`  //更新时间
	Createtime int64  `json:"createtime"` //创建时间
	Page       int64  `json:"page"`       //页码
	PageSize   int64  `json:"page_size"`  //页距
}

//获取表名
func (u *TeacherType) TableName() string {
	return "teacher"
}

//查询内容
func (u *TeacherType) FindOne() (*Teacher, error) {
	list := &Teacher{}
	err := dbs.Table(u.TableName()).Where("id = ?", u.Id).
		Find(list).Error
	if err != nil {
		return &Teacher{}, err
	}
	return list, nil
}

//查询内容
func (u *TeacherType) FindLimit() (*[]Teacher, error) {
	list := &[]Teacher{}
	sqls := dbs.Table(u.TableName())
	sqls = sqls.Order("createtime desc").Limit(u.PageSize).Offset((u.Page - 1) * u.PageSize).Debug()
	err := sqls.Find(list).Error
	if err != nil {
		return &[]Teacher{}, err
	}
	return list, nil
}

//查询所有内容
func (u *TeacherType) FindAll() (*[]Teacher, error) {
	list := &[]Teacher{}
	err := dbs.Table(u.TableName()).
		Find(list).Error
	if err != nil {
		return &[]Teacher{}, err
	}
	return list, nil
}

//删除数据
func (u *TeacherType) Delete() (int, error) {
	err := dbs.Table(u.TableName()).Where("id = ?", u.Id).Delete(&Teacher{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}
