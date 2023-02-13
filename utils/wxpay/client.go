package wxpay

import (
	"context"
	"github.com/google/uuid"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	wxpay "github.com/wechatpay-apiv3/wechatpay-go/utils"
	"gopartsrv/utils/mini"
	"log"
)


//创建支付服务
func createClient() (*core.Client,error) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := wxpay.LoadPrivateKeyWithPath(mini.MchPKFileName)
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mini.Mchid, mini.MchCertificateSerialNumber, mchPrivateKey, mini.MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}
	return client,nil
}

//下单
func CreatOrder(openid string,amount int64) (*jsapi.PrepayWithRequestPaymentResponse,error) {
	client,_:=createClient()
	svc := jsapi.JsapiApiService{Client: client}
	tradeNo := uuid.NewString()[:18]
	resp, _, err := svc.PrepayWithRequestPayment(context.Background(),
		jsapi.PrepayRequest{
			Appid:       core.String(mini.APPID),
			Mchid:       core.String(mini.Mchid),
			Description: core.String(mini.Description),
			OutTradeNo:  core.String(tradeNo),
			Attach:      core.String("龙猫小程序微信支付"),
			NotifyUrl:   core.String(mini.ServerURL),
			Amount: &jsapi.Amount{
				Total: core.Int64(amount),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(openid),
			},
		},
	)
	if err != nil {
		return resp,err
	}
	return resp ,nil
}

