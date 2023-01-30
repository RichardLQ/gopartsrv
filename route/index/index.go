package index

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/controller/index"
)

//首页路由
func IndexRouter(e *gin.Engine) {
	v1 := e.Group("/index")
	{
		v1.GET("/banner", index.Banner)
	}
}
