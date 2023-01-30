package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gopartsrv/condition/model"
	"gopartsrv/public/consts"
	"net/http"
	"time"
)
//ban栏
func Banner(c *gin.Context) {
	pic := &model.PicList{}
	list,err := pic.FindLast()
	if err!=nil {
		fmt.Println("查询失败！")
	}
	old := ""
	if len(*list) > 0 {
		old = time.Unix((*list)[0].Createtime, 0).Format(consts.FORMATDATESHORT)
	}
	now := time.Now().Format(consts.FORMATDATESHORT)
	picList := []*model.PicList{}
	if old != now {
		pic.DeleteAll()
		for i:=0;i<7;i++ {
			temp := consts.HttpGet(consts.PICAPI)
			urls := gjson.Get(temp, "imgurl")
			temps := &model.PicList{
				Imageurl: urls.String(),
				Createtime: time.Now().Unix(),
			}
			picList = append(picList, temps)
		}
		pic.InsertAll(picList)
	}
	list,err = pic.FindAll()
	c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求成功", "data": list})
	return
}
