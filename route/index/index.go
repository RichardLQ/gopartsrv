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
		v1.GET("/order", index.CreatOrder)// 下单
		v1.GET("/payStatus", index.PayStatus)// 支付
		v1.GET("/getOpenid", index.GetOpenid) //获取openid
		v1.GET("/getToken", index.GetTokenTime) //获取accesstoken或者ticket
		v1.GET("/getSign", index.Sign) //获取加密信息



		v1.GET("/getRandomPic", index.GetRandomPic) //获取随机图片
		v1.POST("/orderCallback", index.OrderBack)// 支付回调
	}
}
