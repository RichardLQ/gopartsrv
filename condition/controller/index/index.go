package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/logic/index"
	"net/http"
	"strconv"
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
