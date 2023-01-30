package model

import "gopartsrv/utils/db"

//获取连接
func init() {
	dbs = db.DBTeamMap["work"] //同一个db
}

//banner信息
type Collects struct {
	Id         string `json:"id"`         //id
	Openid     string `json:"openid"`     //openid
	Types      string `json:"types"`      //类型：0:心路历程，1:记忆流沙
	Address    string `json:"address"`    //图片地址
	Collects   string `json:"collects"`   //搜集：0：未搜集，1：搜集
	Updatetime string `json:"updatetime"` //更新时间
	Createtime string `json:"createtime"` //创建时间
}

//获取表名
func (b *Collects) TableName() string {
	return "collect"
}

//查询内容(id 或openid查询)
func (b *Collects) Find() (*[]Collects, error) {
	list := &[]Collects{}
	sqls :=dbs.Table(b.TableName())
	if(b.Id == ""){
		sqls = sqls.Where("id = ?", b.Id)
	}
	if(b.Openid == ""){
		sqls = sqls.Where("openid = ?", b.Openid)
	}
	err := sqls.Find(list).Error
	if err != nil {
		return &[]Collects{}, err
	}
	return list, nil
}

//添加数据
func (u *Collects) Create() error {
	err := dbs.Table(u.TableName()).
		Create(u).Error
	if err != nil {
		return err
	}
	return err
}
