package main

import (
	log "github.com/sirupsen/logrus"
	"gopartsrv/public/service"
	_ "gopartsrv/utils/logs"
)



func main() {
	//var (
	//	mchID                      string = "1637163088"                                // 商户号
	//	mchCertificateSerialNumber string = "4F59EB541378CC84423AE305A596041C776967A1"  // 商户证书序列号
	//	mchAPIv3Key                string = "mADUjnHmfgTtqFk1rPvZJtBcmcqQ6C09"          // 商户APIv3密钥
	//)
	//
	//// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	//mchPrivateKey, err := utils.LoadPrivateKeyWithPath("D:/project/bak/gopartsrv/utils/cert/apiclient_key.pem")
	//if err != nil {
	//	log.Fatal("load merchant private key error")
	//}
	//
	//ctx := context.Background()
	//// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	//opts := []core.ClientOption{
	//	option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	//}
	//client, err := core.NewClient(ctx, opts...)
	//if err != nil {
	//	log.Fatalf("new wechat pay client err:%s", err)
	//}
	//
	//// 发送请求，以下载微信支付平台证书为例
	//// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	//svc := certificates.CertificatesApiService{Client: client}
	//resp, result, err := svc.DownloadCertificates(ctx)
	//log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)


	log.WithFields(log.Fields{"info": "这是golang日志框架--logrus"}).Info("描述信息为golang日志框架logrus的学习")
	service.Serviceinit() //启动服务
}
