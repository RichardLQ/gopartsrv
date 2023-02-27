package mini

import (
	"github.com/tidwall/gjson"
	"gopartsrv/public/consts"
)

const (
	MENU_URL = " https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	TOKEN_URL ="https://sz.api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"
	TICKET_URL ="https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	SIGN_URL = "jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s"
	GRANT_TYPE = "authorization_code"
	//心路历程
	APPID = "wxc055add4d2d04367"
	SECRET = "2c315d992f26cbc89faeab0805b207ab"

	//龙猫公众号
	LMP_APPID="wxde2cf49d6527e57a"
	LMP_SECRET = "b7ed63af9a651915a206ca832a2057e9"
	//下单地址
	ORDER_URL = "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"
	MCHID = "1637163088"
	BF_TOKEN ="SourCandyQiao"
	BF_MSGTYPE = "event" //公众号消息类型
	BF_MSGTYPETEXT = "text" //公众号消息类型
	BF_EVENT_SUBSCEIBE = "subscribe" //订阅

)
//支付参数
const (
	Mchid                      = "1637163088"                                // 商户号
	MchCertificateSerialNumber = "4F59EB541378CC84423AE305A596041C776967A1"  // 商户证书序列号
	MchAPIv3Key                = "mADUjnHmfgTtqFk1rPvZJtBcmcqQ6C09"          // 商户APIv3密钥
	//MchPKFileName			 = "/usr/local/nginx/ssl/apiclient_key.pem"         // 下载的证书文件
	//MchPKFileName			 = "D:/data/conf/apiclient_key.pem"         // 下载的证书文件
	MchPKFileName			 = "G:/data/apiclient_key.pem"         // 下载的证书文件
	ServerURL				 = "https://www.sourcandy.cn/part/index/orderCallback"         // 回调地址
	Description    			 =  "龙猫社群 深圳圈"
)

//获取accesstoken
func GetAccessToken() (token string) {
	filePath := "./accesstoken.txt"
	strs,_ := consts.ReadContent(filePath)
	return gjson.Get(strs, "access_token").String()
}