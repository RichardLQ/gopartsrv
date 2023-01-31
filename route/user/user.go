package user

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/controller/user"
)

//首页路由
func UserRouter(e *gin.Engine) {
	v1 := e.Group("/user")
	{
		v1.GET("/queryLogin", user.UserInfo)
	}
}
