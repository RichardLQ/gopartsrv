package mark

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gopartsrv/public/consts"
	"gopartsrv/utils/mini"
	"log"
	"time"
)

const (
	ROBOT_URL             = "http://api.qingyunke.com/api.php?key=free&appid=0&msg=%s" //机器人地址
	CloudHot_URL          = "https://api.66mz8.com/api/music.163.php"                  //网易云热评
	EpidemicSearch_URL    = "https://szwj.borycloud.com/wh5/index.html#/"              //深圳核酸查询地址
	SERIAL_CloudHot       = "1"                                                        //网易云热评序号
	SERIAL_CallDragon     = "2"                                                        //召唤神龙序号
	SERIAL_EpidemicSearch = "3"                                                        //深圳核酸查询
	Motivational_URL      = "http://api.guaqb.cn/v1/onesaid/"                          //励志名言
	SERIAL_Motivational   = "4"                                                        //励志名言
	SERIAL_EXCEL_FUNC     = "5"                                                        //excel函数视频查询
	SERIAL_EXCEL_FUNC_ADDRESS     = "https://www.bilibili.com/video/BV1Qb4y1o7Bc"       //excel函数地址                                                 //excel函数视频查询
)
//excel函数视频查询
func WxEXCELFUNC(c *gin.Context, textMsg WXTextMsg) {
	repTextMsg := GetWXRepTextMsg(textMsg.ToUserName, textMsg.FromUserName)
	repTextMsg.Content = "<a href='" + SERIAL_EXCEL_FUNC_ADDRESS + "'>深圳核酸查询</a>"
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
	return
}

//励志名言
func WxMotivational(c *gin.Context, textMsg WXTextMsg) {
	str1 := consts.HttpGet(fmt.Sprintf(Motivational_URL, textMsg.Content))
	repTextMsg := GetWXRepTextMsg(textMsg.ToUserName, textMsg.FromUserName)
	repTextMsg.Content = str1
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
	return
}

//深圳疫情查询地址
func WxEpidemicSearch(c *gin.Context, textMsg WXTextMsg) {
	repTextMsg := GetWXRepTextMsg(textMsg.ToUserName, textMsg.FromUserName)
	repTextMsg.Content = "<a href='" + EpidemicSearch_URL + "'>深圳核酸查询</a>"
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
	return
}

//网易云热评
func WxCloudHot(c *gin.Context, textMsg WXTextMsg) {
	repTextMsg := GetWXRepTextMsg(textMsg.ToUserName, textMsg.FromUserName)
	for {
		reStr := consts.HttpGet(CloudHot_URL)
		if gjson.Get(reStr, "comments").String() != "null" {
			content := "  " + gjson.Get(reStr, "comments").String() + "\r\n"
			content += "————网易云音乐热评《" + gjson.Get(reStr, "name").String() + "》"
			repTextMsg.Content = content
			break
		}
	}
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
	return
}

//召唤神龙
func WxCallDragon(c *gin.Context, textMsg WXTextMsg) {
	repTextMsg := GetWXRepTextMsg(textMsg.ToUserName, textMsg.FromUserName)
	repTextMsg.Content = "<a href='https://www.mutegame.com/170/'>召唤神龙</a>"
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
	return
}

//机器人回复
func WxRobot(c *gin.Context, textMsg WXTextMsg) {
	content := "暂时回复不了哦！"
	str1 := consts.HttpGet(fmt.Sprintf(ROBOT_URL, textMsg.Content))
	if gjson.Get(str1, "result").Int() == 0 {
		content = gjson.Get(str1, "content").String()
	}
	repTextMsg := GetWXRepTextMsg(textMsg.ToUserName, textMsg.FromUserName)
	repTextMsg.Content = content
	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
	return
}

//填入回复结构体
func GetWXRepTextMsg(userName, fromName string) WXRepTextMsg {
	repTextMsg := WXRepTextMsg{
		ToUserName:   fromName,
		FromUserName: userName,
		CreateTime:   time.Now().Unix(),
		MsgType:      mini.BF_MSGTYPETEXT,
		Content:      "暂时回复不了哦！",
	}
	return repTextMsg
}
