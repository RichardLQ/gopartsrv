package index

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"gopartsrv/condition/logic/index"
	"gopartsrv/condition/model"
	"gopartsrv/public/consts"
	"gopartsrv/utils/mini"
	"net/http"
	"strconv"
	"time"
)
//Hotlist 热门列表
func Hotlist(c *gin.Context) {
	userId,_ := c.GetPostForm("userid")
	userIds, err := strconv.ParseInt(userId, 10, 64)
	openid,_  := c.GetPostForm("openid")
	list,err:=index.Hotlist(userIds,openid,2)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求失败", "data":list})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data":list})
	return
}


//Partlist 兼职列表
func Partlist(c *gin.Context) {
	userId := c.PostForm("userid")
	openid := c.PostForm("openid")
	pages := c.PostForm("page")
	page, _ := strconv.Atoi(pages)
	pageSizes := c.PostForm("pageSize")
	pageSize, _ := strconv.Atoi(pageSizes)
	search := c.PostForm("search")
	city := c.PostForm("city")
	area := c.PostForm("area")
	list,err:=index.Partlist(userId,openid,search,city,area,page,pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求失败", "data":list})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data":list})
	return
}

//添加
func AddPartlist(c *gin.Context)  {
	status, err := strconv.Atoi(c.PostForm("status"))
	look, err := strconv.Atoi(c.PostForm("look"))
	hot, err := strconv.Atoi(c.PostForm("hot"))
	price, err := strconv.ParseFloat(c.PostForm("price"),64)
	part:= model.Partlist{
		Uid: c.PostForm("uid"),
		Openid:c.PostForm("openid"),
		Status: status,
		Title: c.PostForm("title"),
		Content: c.PostForm("content"),
		Tag: c.PostForm("tag"),
		Price: price,
		Unit: c.PostForm("unit"),
		Province: c.PostForm("province"),
		City: c.PostForm("city"),
		Area: c.PostForm("area"),
		Look: look,
		Hot: hot,
		Createtime: time.Now().Format(consts.FORMATDATELONG),
		Updatetime: time.Now().Format(consts.FORMATDATELONG),
	}
	err = part.Add()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errs": err, "msg": "请求失败", "data":""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errs": "", "msg": "请求成功", "data":""})
	return
}


func GetOpenid(c *gin.Context){
	code := c.Query("code")
	types := c.Query("type")
	urls:=fmt.Sprintf(consts.OPENIDURL,mini.APPID,mini.SECRET,code,mini.GRANT_TYPE)
	if types == "2"{
		urls=fmt.Sprintf(consts.P_OPENIDURL,mini.LMP_APPID,mini.LMP_SECRET,code,mini.GRANT_TYPE)
	}
	fmt.Println(urls)
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

type tokens struct {
	AccessToken string `json:"access_token"`
	ExpireIn int64 `json:"expire_in"`
}

type tickets struct {
	Ticket string `json:"ticket"`
	ExpireIn int64 `json:"expire_in"`
}

func GetTokenTime(c *gin.Context)  {
	types := c.Query("types")
	if types == "token"{
		token:=getToken()
		c.JSON(http.StatusOK, gin.H{"ret":http.StatusOK,"msg": "请求成功", "token": token,"ticket":""})
		return
	}
	token,ticket:=getTicket()
	c.JSON(http.StatusOK, gin.H{"ret":http.StatusOK,"msg": "请求成功", "token": token,"ticket":ticket})
	return
}

//获取accesstoken
func getToken() string {
	filePath := "./accesstoken.txt"
	isExist,token := getTimes(filePath)
	if isExist {
		return token
	}
	grant_type := "client_credential"
	appid := mini.LMP_APPID
	secret := mini.LMP_SECRET
	urls := fmt.Sprintf(mini.TOKEN_URL,grant_type,appid,secret)
	str :=consts.HttpGet(urls)
	content := tokens{
		AccessToken: gjson.Get(str, "access_token").String(),
		ExpireIn: time.Now().Unix(),
	}
	contents, _ := json.Marshal(content)
	consts.WriteContent(filePath,string(contents))
	return gjson.Get(str, "access_token").String()
}

//获取ticket
func getTicket() (string,string) {
	token :=getToken()
	filePath := "./ticket.txt"
	isExist,ticket := getTimes(filePath)
	if isExist {
		return token,ticket
	}
	urls := fmt.Sprintf(mini.TICKET_URL,token)
	str :=consts.HttpGet(urls)
	content := tickets{
		Ticket: gjson.Get(str, "ticket").String(),
		ExpireIn: time.Now().Unix(),
	}
	contents, _ := json.Marshal(content)
	consts.WriteContent(filePath,string(contents))
	return token,gjson.Get(str, "ticket").String()
}

func getTimes(filePath string) (bool,string) {
	content,_:=consts.ReadContent(filePath)
	times := gjson.Get(content,"expire_in").Int()
	token := ""
	isToken := gjson.Get(content,"access_token").Exists()
	isTicket := gjson.Get(content,"ticket").Exists()
	if isToken{
		token = gjson.Get(content,"access_token").String()
	}
	if isTicket{
		token = gjson.Get(content,"ticket").String()
	}
	bend := time.Now().Unix() - times
	if bend < 7100 {
		return true ,token
	}
	return false,""
}

//Sign 加密信息
func Sign(c *gin.Context)  {
	url:=c.Query("url")
	nonceStr := uuid.NewString()[:18]
	_,ticket:= getTicket()
	timestamp := time.Now().Unix()
	urls := fmt.Sprintf(mini.SIGN_URL,ticket,nonceStr,timestamp,url)
	sign := Sha1String(urls)
	data := map[string]interface{}{
		"sign": sign,
		"timestamp":timestamp,
		"nonceStr":nonceStr,
		"appId":mini.LMP_APPID,
	}
	c.JSON(http.StatusOK, gin.H{"ret":http.StatusOK,"msg": "请求成功", "data": data})
	return
}

//sha1加密
func Sha1String( data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}