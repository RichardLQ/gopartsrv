package mini

import (
	"github.com/tidwall/gjson"
	"gopartsrv/public/consts"
)

const (
	MENU_URL = " https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	TOKEN_URL ="https://sz.api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"
	GRANT_TYPE = "authorization_code"
	//心路历程
	APPID = "wxa100ef076707de49"
	SECRET = "b54d634ca6b4f1282417852ca64ae207"
	//记忆流沙
	JY_APPID = "wxa35dbf88d2f73c04"
	JY_SECRET = "13c7dfd67f605328db81ee873cc8fc14"
	//避风港湾
	BF_APPLD = "wxd94dfa3ae4faebe1"
	BF_SECRET ="7e5a0659f5f11dfb6c83a50507ef3381"
	BF_TOKEN ="SourCandyQiao"
	BF_MSGTYPE = "event" //公众号消息类型
	BF_MSGTYPETEXT = "text" //公众号消息类型
	BF_EVENT_SUBSCEIBE = "subscribe" //订阅

)

//获取accesstoken
func GetAccessToken() (token string) {
	filePath := "./accesstoken.txt"
	strs,_ := consts.ReadContent(filePath)
	return gjson.Get(strs, "access_token").String()
}