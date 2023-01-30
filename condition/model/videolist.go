package model

import (
	"gopartsrv/utils/db"
)

//获取连接
func init() {
	dbs = db.DBTeamMap["work"] //同一个db
}

//视频信息
type Videos struct {
	Id         string `json:"id"`         //banner的id
	Videourl   string `json:"videourl"`   //视频地址
	Type       int    `json:"type"`       //类型
	Createtime string `json:"createtime"` //创建时间
	Title      string `json:"title"`      //视频名称
}

//视频结构体
type VideoType struct {
	Id         string `json:"id"`         //banner的id
	Videourl   string `json:"videourl"`   //视频地址
	Type       int    `json:"type"`       //类型
	Createtime string `json:"createtime"` //创建时间
	Title      string `json:"title"`      //视频名称
	Limit      int    `json:"limit"`      //限制
	Page       int64  `json:"page"`
	PageSize   int64  `json:"page_size"`
}

//获取表名
func (u *Videos) TableName() string {
	return "videos"
}

//获取表名
func (u *VideoType) TableName() string {
	return "videos"
}

//查询内容
func (u *Videos) Find() (*Videos, error) {
	list := &Videos{}
	err := dbs.Table(u.TableName()).Where("id = ?", u.Id).
		Find(list).Error
	if err != nil {
		return &Videos{}, err
	}
	return list, nil
}

//根据type查询内容
func (u *Videos) FindType() (*[]Videos, error) {
	list := &[]Videos{}
	err := dbs.Table(u.TableName()).Where("type = ?", u.Type).
		Find(list).Error
	if err != nil {
		return &[]Videos{}, err
	}
	return list, nil
}

//根据type查询内容(限制)
func (u *VideoType) FindTypeLimit() (*[]Videos, error) {
	list := &[]Videos{}
	err := dbs.Table(u.TableName()).Where("type = ?", u.Type).Limit(u.PageSize).Offset((u.Page - 1) * u.PageSize).
		Find(list).Error
	if err != nil {
		return &[]Videos{}, err
	}
	return list, nil
}

//查询所有内容
func (u *Videos) FindAll() (*[]Videos, error) {
	list := &[]Videos{}
	err := dbs.Table(u.TableName()).Limit(4).
		Find(list).Error
	if err != nil {
		return &[]Videos{}, err
	}
	return list, nil
}

//删除数据
func (u *Videos) Delete() (int, error) {
	err := dbs.Table(u.TableName()).Where("id = ?", u.Id).Delete(&Users{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}
