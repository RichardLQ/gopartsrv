package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gopartsrv/condition/logic/index"
	"gopartsrv/condition/model"
	"gopartsrv/public/consts"
	"gopartsrv/utils/mini"
	"net/http"
	"strconv"
	"time"
)
//Hotlist 热门列表
func Hotlist(c *gin.Context) {
	userId,_ := c.GetPostForm("userid")
	openid,_  := c.GetPostForm("openid")
	list,err:=index.Hotlist(userId,openid,2)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求失败", "data":list})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data":list})
	return
}


//Partlist 兼职列表
func Partlist(c *gin.Context) {
	userId := c.PostForm("userid")
	openid := c.PostForm("openid")
	pages := c.PostForm("page")
	page, _ := strconv.Atoi(pages)
	pageSizes := c.PostForm("pageSize")
	pageSize, _ := strconv.Atoi(pageSizes)
	search := c.PostForm("search")
	city := c.PostForm("city")
	area := c.PostForm("area")
	fmt.Println(city)
	fmt.Println(area)
	fmt.Println(search)
	list,err:=index.Partlist(userId,openid,search,city,area,page,pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求失败", "data":list})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data":list})
	return
}


func GetOpenid(c *gin.Context){
	code := c.Query("code")
	urls:=fmt.Sprintf(consts.OPENIDURL,mini.APPID,mini.SECRET,code,mini.GRANT_TYPE)
	respStr:= consts.HttpGet(urls)
	if gjson.Get(respStr,"errcode").Int() != 0{
		c.JSON(http.StatusOK, gin.H{"ret":gjson.Get(respStr,"errcode").Int(),"msg": gjson.Get(respStr,"errmsg").String(), "data": ""})
		return
	}
	user:=model.Users{Openid: gjson.Get(respStr,"openid").String(),Updatetime: time.Now().Format(consts.FORMATDATELONG)}
	ss,_ := user.Find()
	if ss.Id == "" {
		user.Createtime = time.Now().Format(consts.FORMATDATELONG)
		user.Create()
	}else{
		user.Id = ss.Id
		user.Updates()
	}
	c.JSON(http.StatusOK, gin.H{"ret":gjson.Get(respStr,"errcode").Int(),"msg": gjson.Get(respStr,"errmsg").String(), "data": gjson.Get(respStr,"openid").String()})
	return
}