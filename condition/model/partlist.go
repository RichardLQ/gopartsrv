package model

import (
	"gopartsrv/utils/db"
)

// 获取连接
func init() {
	dbs = db.DBTeamMap["mini"] //同一个db
}

// Partlist 订单
type Partlist struct {
	Id         string  `json:"id"` //banner的id
	Openid string `json:"openid"`
	Uid        string  `json:"uid"`
	Status     int     `json:"status"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Tag        string  `json:"tag"`
	Price      float64 `json:"price"`
	Unit       string  `json:"unit"`
	Province   string  `json:"province"`
	City       string  `json:"city"`
	Area       string  `json:"area"`
	Address    string  `json:"address"`
	Look       int     `json:"look"`
	Hot        int     `json:"hot"`
	Buy        bool    `json:"buy"`
	Username string `json:"username"`
	Img string `json:"img"`
	Createtime string  `json:"createtime"` //创建时间
	Updatetime string  `json:"updatetime"`
	Deletetime string  `json:"deletetime"`
}
type PartlistAdd struct {
	Id         string  `json:"id"` //banner的id
	Openid string `json:"openid"`
	Uid        string  `json:"uid"`
	Status     int     `json:"status"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Tag        string  `json:"tag"`
	Price      float64 `json:"price"`
	Unit       string  `json:"unit"`
	Province   string  `json:"province"`
	City       string  `json:"city"`
	Area       string  `json:"area"`
	Address    string  `json:"address"`
	Look       int     `json:"look"`
	Hot        int     `json:"hot"`
	Createtime string  `json:"createtime"` //创建时间
	Updatetime string  `json:"updatetime"`
	Deletetime string  `json:"deletetime"`
}

// 获取表名
func (u *Partlist) TableName() string {
	return "partlist"
}

// 查询内容(id 查询)
func (u *Partlist) Find(limit int, buy bool) (*[]Partlist, error) {
	list := &[]Partlist{}
	sqls := dbs.Debug().Table(u.TableName())
	if u.Hot != 0 {
		sqls = sqls.Where("hot = ?", u.Hot)
	}
	if limit != 0 {
		sqls = sqls.Limit(limit)
	}
	//err := sqls.Joins("left join users on users.id = partlist.uid").Select("*,? as buy,users.address as img", buy).Find(list).Error
	err := sqls.Find(list).Error
	if err != nil {
		return &[]Partlist{}, err
	}
	return list, nil
}

func (u *Partlist) Find2Search(page,pageSize int,buy bool) (*[]Partlist, error) {
	list := &[]Partlist{}
	sqls := dbs.Table(u.TableName())
	if u.Hot != 0 {
		sqls = sqls.Where("hot = ?", u.Hot)
	}
	if u.City != "" {
		sqls = sqls.Where("city = ?", u.City)
	}
	if u.Area != "" {
		sqls = sqls.Where("`area` = ?", u.Area)
	}
	if u.Content != "" {
		sqls = sqls.Where("content like ?", "%"+u.Content+"%")
	}
	//err := sqls.Joins("left join users on users.id = partlist.uid").
	//	Select("*,? as buy,users.address as img", buy).Limit(pageSize).
	//	Offset((page - 1) * pageSize).Find(list).Error
	err := sqls.Offset((page - 1) * pageSize).Find(list).Error
	if err != nil {
		return &[]Partlist{}, err
	}
	return list, nil
}

func (u *Partlist)Add() error {
	err := dbs.Table(u.TableName()).Create(&PartlistAdd{
		Uid: u.Uid,
		Status: u.Status,
		Title: u.Title,
		Content: u.Content,
		Tag: u.Tag,
		Price: u.Price,
		Unit: u.Unit,
		Province: u.Province,
		City: u.City,
		Area: u.Area,
		Address: u.Address,
		Look: u.Look,
		Hot: u.Hot,
		Createtime: u.Createtime,
		Updatetime: u.Updatetime,
		Deletetime: u.Deletetime,
	}).Error
	if err != nil {
		return  err
	}
	return  nil
}