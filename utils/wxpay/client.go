package wxpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"time"
)
var (
	Mchid                      string = "1637163088"                                // 商户号
	MchCertificateSerialNumber string = "4F59EB541378CC84423AE305A596041C776967A1"  // 商户证书序列号
	MchAPIv3Key                string = "mADUjnHmfgTtqFk1rPvZJtBcmcqQ6C09"          // 商户APIv3密钥
	MchPKFileName				string = "D:/project/bak/gopartsrv/utils/cert/apiclient_key.pem"         // 下载的证书文件
	ServerURL					string = "https://www.sourcandy.cn"         // 下载的证书文件
	Appid						string = "wxde2cf49d6527e57a"         // 下载的证书文件
)
//获取加解密处理
func getWechatClient() (context.Context,*core.Client, error) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(MchPKFileName)
	if err != nil {
		log.Print("load merchant private key error")
		return nil,nil, err
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(Mchid, MchCertificateSerialNumber, mchPrivateKey, MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return nil,nil, err
	}
	return ctx,client, nil
}

func CreateWechatPrepayBill(outTradeNo, description, attach, spOpenid string, amount int64) (string, error) {
	notifyUrl := ServerURL +"/api/thirdParty/wechat/pay/getPayResult/"
	ctx,client, err := getWechatClient()
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return "",err
	}

	tmp, _ := time.ParseDuration("5m")
	endTime := time.Now().Add(tmp)
	svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			Appid:         core.String(Appid),
			Mchid:         core.String(Mchid),
			Description:   core.String(description),
			OutTradeNo:    core.String(outTradeNo),
			TimeExpire:    core.Time(endTime),
			Attach:        core.String(attach),
			NotifyUrl:     core.String(notifyUrl),
			GoodsTag:      core.String("WXG"),
			LimitPay:      []string{"no_credit"},
			SupportFapiao: core.Bool(false),
			Amount: &jsapi.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(0.01),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(spOpenid),
			},
			SettleInfo: &jsapi.SettleInfo{
				ProfitSharing: core.Bool(false),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
		return "", nil
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
	return *resp.PrepayId, nil
}


//创建并生成待支付信息
//func CreateWechatPrepayWithPayment(outTradeNo, description, attach, spOpenid string, amount int64)(map[string]interface{}, error){
//	notifyUrl := ServerURL +"/api/thirdParty/wechat/pay/getPayResult"
//	ctx,client, err := getWechatClient()
//	if err != nil {
//		return nil,err
//	}
//
//	tmp, _ := time.ParseDuration("5m")
//	endTime := time.Now().Add(tmp)
//	svc := jsapi.JsapiApiService{Client: client}
//	resp, _, err := svc.PrepayWithRequestPayment(ctx,
//		jsapi.PrepayRequest{
//			Appid:         core.String(Appid),
//			Mchid:         core.String(Mchid),
//			Description:   core.String(description),
//			OutTradeNo:    core.String(outTradeNo),
//			TimeExpire:    core.Time(endTime),
//			Attach:        core.String(attach),
//			NotifyUrl:     core.String(notifyUrl),
//			GoodsTag:      core.String("WXG"),
//			LimitPay:      []string{"no_credit"},
//			SupportFapiao: core.Bool(false),
//			Amount: &jsapi.Amount{
//				Currency: core.String("CNY"),
//				Total:    core.Int64(0.01),
//			},
//			Payer: &jsapi.Payer{
//				Openid: core.String(spOpenid),
//			},
//			SettleInfo: &jsapi.SettleInfo{
//				ProfitSharing: core.Bool(false),
//			},
//		},
//	)
//
//	if err != nil {
//		// 处理错误
//		return nil, err
//	} else {
//		// 处理返回结果
//		//log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
//	}
//	result := make(map[string]interface{})
//	tmpJson := utils.GetJsonStr(resp)
//	json.Unmarshal([]byte(tmpJson),&result)
//
//	return result, nil
//}