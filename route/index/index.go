package index

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/controller/index"
)

//首页路由
func IndexRouter(e *gin.Engine) {
	v1 := e.Group("/index")
	{
		v1.POST("/hotlist", index.Hotlist)// 热门列表
		v1.POST("/partlist", index.Partlist)// 兼职列表
	}
}
