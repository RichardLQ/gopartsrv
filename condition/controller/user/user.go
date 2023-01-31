package user

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/logic/user"
	"gopartsrv/public/consts"
	"log"
	"net/http"
)

//用户信息
func UserInfo(c *gin.Context) {
	userId := c.Query("userid")
	openid := c.Query("openid")
	log.Printf("UserInfo %v; %v \n", userId, openid)
	if userId == "" && openid == "" {
		c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "请求错误", "data": "", "code": consts.PARAM_ERROR})
		return
	}
	list, err := user.UserInfo(userId,openid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求错误", "data": "", "code": consts.SEARCH_FAIL})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": list, "code": http.StatusOK})
	return
}

