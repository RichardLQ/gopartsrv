package index

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/logic/index"
	"net/http"
	"strconv"
)
//Hotlist 热门列表
func Hotlist(c *gin.Context) {
	userId := c.Query("userid")
	openid := c.Query("openid")
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
	userId := c.Query("userid")
	openid := c.Query("openid")
	pages := c.Query("page")
	page, _ := strconv.Atoi(pages)
	pageSizes := c.Query("pageSize")
	pageSize, _ := strconv.Atoi(pageSizes)
	search := c.Query("search")
	city := c.Query("city")
	area := c.Query("area")
	list,err:=index.Partlist(userId,openid,search,city,area,page,pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求失败", "data":list})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data":list})
	return
}
