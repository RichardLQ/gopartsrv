package model

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopartsrv/utils/db"
)

var (
	sqlites *gorm.DB
)

//获取连接
func init() {
	sqlites = db.SqliteTeamMap["piclist"] //同一个db
}

type PicList struct {
	Id         string    `json:"id"`
	Imageurl   string `json:"imageurl"`
	Createtime int64  `json:"createtime"`
}

//获取表名
func (u *PicList) TableName() string {
	return "piclist"
}

//查询内容
func (u *PicList) Find() (*PicList, error) {
	list := &PicList{}
	err := sqlites.Table(u.TableName()).Where("id = ?", u.Id).
		Find(list).Error
	if err != nil {
		return &PicList{}, err
	}
	return list, nil
}
//查询最后一个内容
func (u *PicList) FindLast() (*[]PicList, error) {
	list := &[]PicList{}
	err := sqlites.Table(u.TableName()).
		Last(list).Error
	if err != nil {
		return &[]PicList{}, err
	}
	return list, nil
}

//查询内容
func (u *PicList) FindAll() (*[]PicList, error) {
	list := &[]PicList{}
	err := sqlites.Table(u.TableName()).
		Find(list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

//添加数据
func (u *PicList) Create() error {
	err := dbs.Table(u.TableName()).
		Create(u).Error
	if err != nil {
		return err
	}
	return err
}

//删除数据
func (u *PicList) Delete() (int, error) {
	err := sqlites.Table(u.TableName()).Where("id = ?", u.Id).Delete(&PicList{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}
//删除所有数据
func (u *PicList) DeleteAll() (int, error) {
	err := sqlites.Table(u.TableName()).Where("1 = 1").Delete(&PicList{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

//批量插入
func (u *PicList) InsertAll(emps []*PicList) error {
	var buffer bytes.Buffer
	sql := "insert into " +u.TableName()+ " (`imageurl`,`createtime`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range emps {
		if i == len(emps)-1 {
			buffer.WriteString(fmt.Sprintf("('%s',%d);", e.Imageurl, e.Createtime))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s',%d),",e.Imageurl, e.Createtime))
		}
	}
	return sqlites.Exec(buffer.String()).Error
}