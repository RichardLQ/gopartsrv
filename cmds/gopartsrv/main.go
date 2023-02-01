package main

import (
	log "github.com/sirupsen/logrus"
	"gopartsrv/public/consts"
	"gopartsrv/public/service"
	_ "gopartsrv/utils/logs"
)



func main() {
	//var (
	//	mchID                      string = "1637163088"                                // 商户号
	//	mchCertificateSerialNumber string = "4F59EB541378CC84423AE305A596041C776967A1"  // 商户证书序列号
	//	mchAPIv3Key                string = "mADUjnHmfgTtqFk1rPvZJtBcmcqQ6C09"          // 商户APIv3密钥
	//)
	url := "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"
	data := `{
	"mchid": "1637163088",
	"out_trade_no": "1217752501201407033233368318",
	"appid": "wxdace645e0bc2cXXX",
	"description": "Image形象店-深圳腾大-QQ公仔",
	"notify_url": "https://www.weixin.qq.com/wxpay/pay.php",
	"amount": {
		"total": 1,
		"currency": "CNY"
	},
	"payer": {
		"openid": "o4GgauInH_RCEdvrrNGrntXDuXXX"
	}
}`
	res := consts.HttpPost(url, data, "")

	log.Printf("res: %v\n", res)
	log.WithFields(log.Fields{"info": "这是golang日志框架--logrus"}).Info("描述信息为golang日志框架logrus的学习")
	service.Serviceinit() //启动服务
}
