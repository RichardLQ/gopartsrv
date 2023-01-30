package user

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/logic/user"
	"gopartsrv/public/consts"
	"net/http"
)

//用户信息
func UserInfo(c *gin.Context) {
	userId := c.Query("userid")
	if userId == "" {
		c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "请求错误", "data": "", "code": consts.PARAM_ERROR})
		return
	}
	list, err := user.UserInfo(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求错误", "data": "", "code": consts.SEARCH_FAIL})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": list, "code": http.StatusOK})
	return
}

func Selection(c *gin.Context)  {
	
}
