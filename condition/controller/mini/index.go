package mini

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/model"
	"gopartsrv/public/config"
	"gopartsrv/public/consts"
	"gopartsrv/utils/qiniu"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetPicUrl1(c *gin.Context){
	imgPath := "./public/images/"
	imgUrl := c.Query("imgUrl")
	fileName := consts.Uuid()+".jpg"
	res, err := http.Get(imgUrl)
	if err != nil {
		return
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)
	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	io.Copy(writer, reader)
	c.JSON(http.StatusOK, gin.H{"ret":0,"msg": "请求成功！", "data": fileName})
	return
}

func GetPicUrl(c *gin.Context){
	str := c.Query("imgUrl")
	strs := strings.Split(str,"/")
	imgUrl:=strings.Split(strs[3],"?")[0]
	code ,_:=http.Get("https://cdn.sourcandy.cn/uploads/"+imgUrl+".jpg")
	if code.StatusCode == 404 {
		address,filename,_,err := qiniu.QiNiu_ByteUploadFile(str,"uploads",imgUrl)
		if err !=nil {
			c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求错误", "data": "", "code": consts.SEARCH_FAIL})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ret":0,"msg": "请求成功！", "data": map[string]string{"address":address,"filename":filename}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ret":0,"msg": "请求成功！", "data": map[string]string{"address":"https://cdn.sourcandy.cn/uploads/"+imgUrl+".jpg","filename":imgUrl+".jpg"}})
	return

}

//视频内容
func Review(c *gin.Context)  {
	videos :=model.Videos{}
	list, err := videos.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求错误", "data": "", "code": consts.SEARCH_FAIL})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": list, "code": http.StatusOK})
	return
}

//公司信息
func CompanyInfo(c *gin.Context)  {
	isType := config.GetIsType()
	data :=map[string]interface{}{
		"isType":isType,
		"companyName":"LINK（杭州）舞蹈CLUB",
		"businessHours":"周一至周日 早10:00-22:00",
		"address":"杭州市拱墅区万达广场3F",
		"phone":"15700177960",
		"scopeBusiness":"舞蹈是人类历史上最早产生的艺术形式之一，人们称之为“艺术之母”，它随着历史的进步而变化发展。",
		"introduction":"舞蹈是人类历史上最早产生的艺术形式之一，人们称之为“艺术之母”，它随着历史的进步而变化发展。",
		"imageurl":"",
		"videourl":"",
		"videoTitle":"",
	}
	types := c.DefaultQuery("type","3")
	banners :=model.Banners{Type: types,Ext: "limit"}
	list, err := banners.FindType()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求错误", "data": "", "code": consts.SEARCH_FAIL})
		return
	}
	data["imageurl"] = (*list)[0].Imageurl
	vi := model.Videos{
		Type: 5,
	}
	if isType == 2 {
		video , err:= vi.FindType()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求错误", "data": "", "code": consts.SEARCH_FAIL})
			return
		}
		data["videourl"] = (*video)[0].Videourl
		data["videoTitle"] = (*video)[0].Title
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": data, "code": http.StatusOK})
	return
}

//热门视频和图片
func GetHotImageAndVideo(c *gin.Context)  {
	types := config.GetIsType()
	if types == 1 {
		ban := model.Banners{
			Type: "4",
		}
		list,_ := ban.FindType()
		resp := map[string]interface{}{
			"list":list,"type":types,
		}
		c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": resp, "code": http.StatusOK})
		return
	}
	if types == 2 {
		video := model.Videos{}
		list,_ := video.FindAll()
		resp := map[string]interface{}{
			"list":list,"type":types,
		}
		c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": resp, "code": http.StatusOK})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "没有找到！", "data": "", "code": http.StatusOK})
	return
}

//舞蹈类型内容
func GetTypeVideo(c *gin.Context)  {
	typ := c.Query("type")
	page := c.DefaultQuery("page","1")
	pages, _ := strconv.ParseInt(page, 10, 64)
	pageSize := c.DefaultQuery("pageSize","4")
	pageSizes, _ := strconv.ParseInt(pageSize, 10, 64)
	if typ == "" {
		c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "参数缺失！", "data": "", "code": 0})
		return
	}
	types := config.GetIsType()
	if types == 1 {
		ban := model.Banners{
			Type: "4",
		}
		if pages == 2 {
			c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": map[string]interface{}{"list":[]model.Banners{},"type":types}, "code": http.StatusOK})
			return
		}
		list,_ := ban.FindType()
		resp := map[string]interface{}{
			"list":list,"type":types,
		}
		c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": resp, "code": http.StatusOK})
		return
	}
	types1, _ := strconv.Atoi(typ)
	vi := model.VideoType{
		Type:types1,
		Limit: 4,
		Page: pages,
		PageSize: pageSizes,
	}
	list,_ := vi.FindTypeLimit()
	resp := map[string]interface{}{
		"list":list,"type":types,
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": resp, "code": http.StatusOK})
	return
}

//获取老师信息
func GetTeacherInfo(c *gin.Context)  {
	page := c.DefaultQuery("page","1")
	pages, _ := strconv.ParseInt(page, 10, 64)
	pageSize := c.DefaultQuery("pageSize","4")
	pageSizes, _ := strconv.ParseInt(pageSize, 10, 64)
	teach := model.TeacherType{
		Page: pages,
		PageSize: pageSizes,
	}
	list,err := teach.FindLimit()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求失败！", "data": "", "code": 0})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": list, "code": http.StatusOK})
	return
}
