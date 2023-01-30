package middle

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//token验证必须登录
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); !status {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    1,
				"message": "必须传递token",
			})
			c.Abort()
			return
		}
	}
}
