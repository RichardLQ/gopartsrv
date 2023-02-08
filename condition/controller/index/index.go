package index

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/logic/index"
	"net/http"
)
//Hotlist 热门列表
func Hotlist(c *gin.Context) {
	userId := c.Query("userid")
	openid := c.Query("openid")
	list,_:=index.Hotlist(userId,openid,2)
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data":list})
	return
}
