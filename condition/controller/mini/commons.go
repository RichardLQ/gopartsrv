package mini

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/tidwall/gjson"
	"gopartsrv/condition/model"
	"gopartsrv/public/consts"
	"gopartsrv/utils/mini"
	"gopartsrv/utils/qiniu"
	"net/http"
	"strconv"
	"strings"
	"time"
)
type MyClams struct {
	UserName string `json:"username"`
	Password string  `json:"password"`
	jwt.StandardClaims
}
//首页banner
func Banner(c *gin.Context){
	types := c.DefaultQuery("type","0")
	banners :=model.Banners{Type: types}
	list, err := banners.FindType()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求错误", "data": "", "code": consts.SEARCH_FAIL})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "返回成功！", "data": list, "code": http.StatusOK})
	return
}

//获取openid
func GetOpenid(c *gin.Context) {
	code := c.Query("code")
	urls:=fmt.Sprintf(consts.OPENIDURL,mini.APPID,mini.SECRET,code,mini.GRANT_TYPE)
	respStr:= consts.HttpGet(urls)
	if gjson.Get(respStr,"errcode").Int() != 0{
		c.JSON(http.StatusOK, gin.H{"ret":gjson.Get(respStr,"errcode").Int(),"msg": gjson.Get(respStr,"errmsg").String(), "data": ""})
		return
	}
	user:=model.Users{Openid: gjson.Get(respStr,"openid").String(),Updatetime: time.Now().Format(consts.FORMATDATELONG)}
	ss,_ := user.Find()
	if ss.Id == "" {
		user.Createtime = time.Now().Format(consts.FORMATDATELONG)
		user.Create()
	}else{
		user.Id = ss.Id
		user.Updates()
	}
	c.JSON(http.StatusOK, gin.H{"ret":gjson.Get(respStr,"errcode").Int(),"msg": gjson.Get(respStr,"errmsg").String(), "data": gjson.Get(respStr,"openid").String()})
	return
}

//获取token
func GetToken(c *gin.Context)  {
	//Keys := []byte("1234565")
	//me := MyClams{
	//	UserName: "Jerome",
	//	Password: "imslimsl",
	//	StandardClaims:jwt.StandardClaims{
	//		NotBefore: time.Now().Unix()-60,
	//		ExpiresAt: time.Now().Unix()-60*60*2,
	//		Issuer: "Jerome",
	//	},
	//}
	//	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,me)
	//	ss,err:= token.SignedString(Keys)

}

//执行脚本
func GetScriptCmd(c *gin.Context)  {
	urls := c.Query("urls")
	params := []string{
		"./test.py",urls,
	}
	fmt.Println()
	str,err:=consts.CmdScript("python",params)
	c.JSON(http.StatusOK, gin.H{"ret":"成功","data":str,"err":err})
	return
}

//更新用户信息
func UpdateUserInfo(c *gin.Context)  {
	avatarUrl := c.Query("avatarUrl")
	openid := c.Query("openid")
	nickName := c.Query("nickName")
	gender := c.Query("gender")
	user:=model.Users{
		Address: avatarUrl,
		Nickname: nickName,
		Openid: openid,
		Gender: gender,
		Updatetime: time.Now().Format(consts.FORMATDATELONG),
	}
	count,err := user.Updates()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"ret":"失败","data":"","err":err,"code":http.StatusIMUsed})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ret":"成功","data":count,"err":err,"code":http.StatusOK})
}

//搜集
func AddCollectInfo(c *gin.Context)  {
	openid := c.Query("openid")
	types := c.Query("type")
	address := c.Query("address")
	collects := c.Query("collects")
	collect := model.Collects{
		Openid: openid,
		Types: types,
		Address: address,
		Collects: collects,
		Updatetime: time.Now().Format(consts.FORMATDATELONG),
		Createtime: time.Now().Format(consts.FORMATDATELONG),
	}
	c.JSON(http.StatusOK, gin.H{"ret":"成功","data":collect,"err":"","code":http.StatusOK})
	return
}

//上传图片
func UploadImage(c *gin.Context)  {
	file, err := c.FormFile("file")
	openid := c.PostForm("openid")
	types,_:=strconv.Atoi(c.PostForm("types"))
	if openid == ""  {
		c.JSON(http.StatusOK, gin.H{"ret":"查询失败","data":"","err":"参数缺失","code":consts.PARAM_LACK})
		return
	}
	if err != nil {
		c.String(500, "上传图片出错")
	}
	str1,_,_,err:=qiniu.QiNiu_SourceUploadFile(file,"load","")
	if err !=nil {
		c.JSON(http.StatusOK, gin.H{"ret":"上传失败","data":"","err":err,"code":-1})
		return
	}
	if str1 != "" {
		pic := model.UploadFile{
			Openid: openid,
			Address: str1,
			Type: types,
			Createtime: time.Now().Format(consts.FORMATDATELONG),
		}
		err:=pic.Create()
		if err !=nil {
			c.JSON(http.StatusOK, gin.H{"ret":"上传失败","data":"","err":err,"code":-2})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"ret":"成功","data":str1,"err":"","code":http.StatusOK})
	return
}

//查看上传图片
func GetUploadImage(c *gin.Context)  {
	openid := c.Query("openid")
	types,_:=strconv.Atoi(c.Query("types"))
	page,_:=strconv.Atoi(c.DefaultQuery("page","1"))
	pageSize,_:=strconv.Atoi(c.DefaultQuery("pageSize","10"))
	if openid == "" {
		c.JSON(http.StatusOK, gin.H{"ret":"查询失败","data":"","err":"参数缺失","code":consts.PARAM_LACK})
		return
	}
	pic := model.UploadFile{
		Openid: openid,
		Type: types,
		Page: page,
		PageSize: pageSize,
	}
	list,err := pic.FindArray()
	if err !=nil {
		c.JSON(http.StatusOK, gin.H{"ret":"查询失败","data":"","err":err,"code":-1})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ret":"成功","data":list,"err":"","code":http.StatusOK})
	return
}

//删除照片
func DelUploadImage(c *gin.Context)  {
	image_id:= c.Query("image_id")
	address:= c.Query("address")
	if image_id == "" || address == ""{
		c.JSON(http.StatusOK, gin.H{"ret":"删除失败","data":"","err":"参数缺失","code":consts.PARAM_LACK})
		return
	}
	addressList := strings.Split(address,"/")
	adder := addressList[len(addressList)-2] + "/" + addressList[len(addressList)-1]
	fmt.Println(adder)
	err:=qiniu.QiNiu_DeleteFile(adder)
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{"ret":"删除失败","data":"","err":err,"code":-2})
		return
	}
	pic := model.UploadFile{
		Id: image_id,
	}
	list,err := pic.Delete()
	if err !=nil {
		c.JSON(http.StatusOK, gin.H{"ret":"删除失败","data":"","err":err,"code":-1})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ret":"删除成功","data":list,"err":"","code":http.StatusOK})
	return
}























