package model

import (
	"gopartsrv/utils/db"
)

//获取连接
func init() {
	dbs = db.DBTeamMap["work"] //同一个db
}

//banner信息
type Banners struct {
	Id         string `json:"id"` //banner的id
	Imageurl   string `json:"imageurl"` //图片地址
	Title   string `json:"title"` //图片标题
	Createtime   string `json:"createtime"` //创建时间
	Type   string `json:"type"` //类型：0:心路历程，1:记忆流沙
	Ext    string `json:"ext"` //额外增加字段
}
//获取表名
func (b *Banners) TableName() string {
	return "banner"
}

//查询内容(id 查询)
func (b *Banners) Find() (*Banners, error) {
	list := &Banners{}
	err := dbs.Table(b.TableName()).Where("id = ?", b.Id).
		Find(list).Error
	if err != nil {
		return &Banners{}, err
	}
	return list, nil
}

//查询内容(type 查询)
func (b *Banners) FindLimit() (*[]Banners, error){
	list := &[]Banners{}
	sqls :=dbs.Table(b.TableName())
	if(b.Ext =="limit"){
		sqls = sqls.Limit(4)
	}
	err := sqls.Where("type = ?", b.Type).
		Find(list).Error
	if err != nil {
		return &[]Banners{}, err
	}
	return list, nil
}

//查询内容(type 查询)
func (b *Banners) FindType() (*[]Banners, error){
	list := &[]Banners{}
	sqls :=dbs.Table(b.TableName())
	if(b.Ext =="limit"){
		sqls = sqls.Limit(1)
	}
	err := sqls.Where("type = ?", b.Type).
		Find(list).Error
	if err != nil {
		return &[]Banners{}, err
	}
	return list, nil
}

//删除数据
func (b *Banners) Delete() (int, error) {
	err := dbs.Table(b.TableName()).Where("id = ?", b.Id).Delete(&Banners{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}