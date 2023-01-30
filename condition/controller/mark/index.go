package mark

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"gopartsrv/public/consts"
	"gopartsrv/utils/db"
	"gopartsrv/utils/mini"
	"log"
	"sort"
	"strings"
)

//接受消息体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	Event        string
	MsgId        int64
}
//发送消息体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

//验证服务器
func GetSignire(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	var tempArray = []string{mini.BF_TOKEN, timestamp, nonce}
	sort.Strings(tempArray)
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	//获得加密后的字符串可与signature对比
	if sha1String == signature {
		c.Writer.Write([]byte(echostr))
	} else {
		log.Println("微信API验证失败")
	}
}

//创建菜单
func CreateMenu(c *gin.Context) {
	menuStr := map[string]interface{}{
		"button": []map[string]interface{}{
			{
				"name": "进入商城",
				"type": "view",
				"url":  "http://www.baidu.com/",
			},
			{
				"name": "管理中心",
				"type": "view",
				"url":  "http://www.baidu.com/",
			},
			{
				"name": "资料修改",
				"type": "view",
				"url":  "http://www.baidu.com/user_view",
			},
		},
	}
	consts.HttpPost(fmt.Sprintf(mini.MENU_URL, mini.GetAccessToken()), menuStr, "")
}

//接受消息
func WxRevice(c *gin.Context) {
	var textMsg WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}
	// 对接收的消息进行被动回复
	wxMsgType(c, textMsg)
	//WxRobot(c, textMsg)

}

//消息类型
func wxMsgType(c *gin.Context, textMsg WXTextMsg) {
	redisCon := db.RedisPoolMap["work"].Get()
	defer redisCon.Close()
	//推送事件
	if textMsg.MsgType == mini.BF_MSGTYPE {
		if textMsg.Event == mini.BF_EVENT_SUBSCEIBE {
			WXMsgSubscribeReply(c,textMsg)
			return
		}
	}
	//普通文本事件
	if textMsg.MsgType == mini.BF_MSGTYPETEXT {
		str := strings.Trim(textMsg.Content," ")

		if str == "主页菜单" {
			_,err := redisCon.Do("SET",textMsg.FromUserName,"yes","EX","7200")
			if err != nil{
				log.Fatal("主页菜单设置失败，",err)
			}
			WXMsgSubscribeReply(c,textMsg)
			return
		}
		if str == "退出菜单" {
			redisCon.Do("DEL",textMsg.FromUserName)
			log.Println(333)
			textMsg.Content = "退出菜单成功！"
			WXMsgReply(c,textMsg)
			return
		}
		redisUserOpenid, _ := redis.String(redisCon.Do("GET",textMsg.FromUserName))
		if redisUserOpenid == "yes" {
			switch str {
			case SERIAL_CloudHot://网易云
				WxCloudHot(c,textMsg)
				return
			case SERIAL_CallDragon://召唤神龙
				WxCallDragon(c,textMsg)
				return
			case SERIAL_EpidemicSearch://深圳核酸查询
				WxEpidemicSearch(c,textMsg)
				return
			case SERIAL_Motivational://励志名言
				WxMotivational(c,textMsg)
				return
			case SERIAL_EXCEL_FUNC://excel函数
				WxEXCELFUNC(c,textMsg)
				return
			default://展示菜单
				WXMsgSubscribeReply(c,textMsg)
				return
			}
		}
		WxRobot(c, textMsg)
		return
	}

	WxRobot(c, textMsg)
	return
}

//关注回复
func WXMsgSubscribeReply(c *gin.Context, textMsg WXTextMsg) {
	content := "《主页菜单》,进入菜单！！！\n"
	content += " " + SERIAL_CloudHot+".网易云热评\n"
	content += " " + SERIAL_CallDragon+".召唤神龙\n"
	content += " " + SERIAL_EpidemicSearch+".深圳核酸检测查询\n"
	content += " " + SERIAL_Motivational+".励志名言\n"
	content += " " + SERIAL_EXCEL_FUNC+".cxcel函数\n"
	content += "《退出菜单》，退出菜单，进入闲聊！！！"
	repTextMsg := GetWXRepTextMsg(textMsg.ToUserName,textMsg.FromUserName)
	repTextMsg.Content = content
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
	return
}

// WXMsgReply 微信消息回复
func WXMsgReply(c *gin.Context,textMsg WXTextMsg) {
	repTextMsg := GetWXRepTextMsg(textMsg.ToUserName,textMsg.FromUserName)
	repTextMsg.Content = textMsg.Content
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复0] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
	return
}
