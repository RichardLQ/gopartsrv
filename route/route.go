package route

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/route/index"
	"gopartsrv/route/mini"
	"gopartsrv/route/user"
	"net/http"
)

func RouteInit(e *gin.Engine) {
	e.StaticFS("/static", http.Dir("./public/images"))
	index.IndexRouter(e)
	mini.MiniRouter(e)
	user.UserRouter(e)
}
