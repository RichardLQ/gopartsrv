package mini

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gopartsrv/condition/model"
	"gopartsrv/public/config"
	"gopartsrv/public/consts"
	"net/http"
	"strconv"
	"time"
)

const API_HOST = "https://service.picasso.adesk.com/"

//热门推荐
func Recommend(c *gin.Context) {
	limit := c.Query("limit")
	skip := c.Query("skip")
	if config.GetIsType() == 1 {
		skip = "1"
	}
	urls := fmt.Sprintf(API_HOST+"v1/vertical/vertical?adult=false&first=20&order=new&limit=%s&skip=%s",limit,skip)
	restr:=consts.HttpGet(urls)
	if gjson.Get(restr,"code").Int() != 0 {
		c.JSON(http.StatusOK, gin.H{"ret":consts.REQUEST_INTERFACE,"msg": gjson.Get(restr,"msg").String(), "data": ""})
		return
	}
	value := gjson.Get(restr,"res.vertical").Array()
	returns := []interface{}{}
	for _,item := range value {
		returns = append(returns, map[string]string{"wp":gjson.Get(item.Raw,"wp").String(),"img":gjson.Get(item.Raw,"wp").String()})
	}
	c.JSON(http.StatusOK, gin.H{"ret":gjson.Get(restr,"code").Int(),"msg": "请求成功", "data": returns})
	return
}

//分类
func Classification(c *gin.Context){
	urls := API_HOST + "v1/vertical/category"
	restr:=consts.HttpGet(urls)
	if gjson.Get(restr,"code").Int() != 0 {
		c.JSON(http.StatusOK, gin.H{"ret":consts.REQUEST_INTERFACE,"msg": gjson.Get(restr,"msg").String(), "data": ""})
		return
	}
	value := gjson.Get(restr,"res.category").Array()
	returns := map[int]interface{}{}
	for index,item := range value {
		returns[index] = map[string]string{"id":gjson.Get(item.Raw,"id").String(),"rname":gjson.Get(item.Raw,"rname").String()}
	}
	c.JSON(http.StatusOK, gin.H{"ret":gjson.Get(restr,"code").Int(),"msg": "请求成功", "data": returns})
	return
}

//分类内容
func ClassifiedContent(c *gin.Context) {
	category := c.Query("category")
	limit := c.Query("limit")
	skip := c.Query("skip")
	if config.GetIsType() == 1 {
		category = "5109e04e48d5b9364ae9ac45"
	}
	urls := fmt.Sprintf(API_HOST+"v1/vertical/category/%s/vertical?adult=false&first=1&order=hot&limit=%s&skip=%s",category,limit,skip)
	restr:=consts.HttpGet(urls)
	if gjson.Get(restr,"code").Int() != 0 {
		c.JSON(http.StatusOK, gin.H{"ret":consts.REQUEST_INTERFACE,"msg": gjson.Get(restr,"msg").String(), "data": ""})
		return
	}
	value := gjson.Get(restr,"res.vertical").Array()
	returns := []interface{}{}
	for _,item := range value {
		returns = append(returns, map[string]string{"wp":gjson.Get(item.Raw,"wp").String(),"img":gjson.Get(item.Raw,"wp").String()})
	}
	c.JSON(http.StatusOK, gin.H{"ret":gjson.Get(restr,"code").Int(),"msg": "请求成功", "data": returns})
	return
}

//投诉建议
func UpdateComplaints(c *gin.Context)  {
	openid := c.Query("openid")
	content := c.Query("content")
	if openid == "" || content == "" {
		c.JSON(http.StatusOK, gin.H{"ret":consts.PARAM_LACK,"msg": "请求失败", "data": "","err":"参数缺失"})
		return
	}
	plaints := model.Complaints{
		Openid: openid,
		Content: content,
		Createtime: time.Now().Format(consts.FORMATDATELONG),
	}
	err:=plaints.Create()
	c.JSON(http.StatusOK, gin.H{"ret":http.StatusOK,"msg": "请求成功", "data": "","err":err})
	return
}
//搜索
func GetSearch(c *gin.Context)  {
	kw := c.Query("kw")
	if config.GetIsType() == 1 {
		kw = "精美壁纸"
	}
	starts := c.Query("start")
	urls := "https://www.duitang.com/napi/blogv2/list/by_search/?kw="+kw+"&after_id="+ starts + "&type=feed&include_fields=top_comments%2Cis_root%2Csource_link%2Citem%2Cbuyable%2Croot_id%2Cstatus%2Clike_count%2Clike_id%2Csender%2Calbum%2Creply_count%2Cfavorite_blog_id&_type=&_=1639233452849"
	res :=consts.HttpGet(urls)
	c.JSON(http.StatusOK, gin.H{"ret":http.StatusOK,"msg": "请求成功", "data": gjson.Get(res,"data").Map(),"err":starts})
	return
}

//添加记事本
func AddNoteContent(c *gin.Context)  {
	openid := c.Query("openid")
	content := c.Query("content")
	createtime := c.Query("createtime")
	if content == "" {
		c.JSON(http.StatusOK, gin.H{"ret":0,"msg": "内容不能为空", "data": ""})
		return
	}
	node := model.Note{
		Openid: openid,
		Createtime: createtime,
		Content: content,
		Updatetime:time.Now().Format(consts.FORMATDATELONG),
	}
	rep,_ := node.Find()
	if rep.Id == 0{
		res,_:= node.Create()
		c.JSON(http.StatusOK, gin.H{"ret":200,"msg": "添加成功", "data": res})
		return
	}else{
		node.Id = rep.Id
		res,_:=node.Update()
		c.JSON(http.StatusOK, gin.H{"ret":200,"msg": "更新成功", "data": res})
		return
	}
}

//删除记事本
func DelNoteContent(c *gin.Context)  {
	index := c.Query("index")
	indexs, err := strconv.ParseInt(index,10,64)
	if index == ""{
		c.JSON(http.StatusOK, gin.H{"ret":0,"msg": "参数缺失", "data": "","err":err})
		return
	}
	node := model.Note{
		Id: indexs,
	}
	res,err:=node.Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"ret":0,"msg": "删除失败", "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ret":200,"msg": "删除成功", "data": res})
	return
}

//查询记事本
func GetNoteContent(c *gin.Context)  {
	openid := c.Query("openid")
	if openid == ""{
		c.JSON(http.StatusOK, gin.H{"ret":0,"msg": "参数缺失", "data": ""})
		return
	}
	page := c.DefaultQuery("page","1")
	pages, err := strconv.ParseInt(page, 10, 64)
	pageSize := c.DefaultQuery("pageSize","10")
	pageSizes, err := strconv.ParseInt(pageSize, 10, 64)
	node := model.Note{
		Openid: openid,
	}
	list ,err :=node.FindArray(pages,pageSizes)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"ret":0,"msg": "查询失败", "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ret":200,"msg": "请求成功", "data": list})
	return
}

